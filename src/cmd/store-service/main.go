package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"frontend"

	"github.com/braintree/manners"
	"github.com/caihua-yin/go-common/common"
)

func main() {
	// Initialize common startup setting
	common.Startup()

	// Load configs
	config := common.LoadGlobConfigs("config.*.yaml")
	// common.MergeConfigFromConsul(config)

	// Initialize store handler
	storeHandler, err := frontend.NewStoreHandler()
	if err != nil {
		log.Fatalf("Initialize store handler failed: %s", err)
	}

	apiMux := http.NewServeMux()
	apiMux.Handle("/store/", storeHandler)
	apiHandler := http.Handler(apiMux)

	rootMux := http.NewServeMux()
	rootMux.Handle("/_health", frontend.NewHealthCheckHandler())
	rootMux.Handle("/_version", frontend.NewVersionHandler())
	rootMux.Handle("/", apiHandler)

	log.Printf("Starting HTTP listener on %s...", config.GetString("http.endpoint"))
	httpServer := manners.NewWithServer(
		&http.Server{
			Addr:    config.GetString("http.endpoint"),
			Handler: rootMux,
		})

	// Listen for signals and do graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigCh

		log.Printf("Shutting down...")

		// Stop catching signals, so that we can actually stop on second signal
		signal.Stop(sigCh)

		atomic.StoreInt32(&frontend.Shutdown, 1)

		// Slow down shutdown to get some extra time for graceful restart
		log.Printf("Sleeping...")
		time.Sleep(common.ParseDuration(config.GetString("shutdown.sleep")))
		log.Printf("Done sleeping!")

		// Shutdown HTTP
		httpServer.Close()
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := httpServer.ListenAndServe()

		if err != nil {
			log.Fatalf("Listen error: %s", err)
		}
	}()

	wg.Wait()
}
