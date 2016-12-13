package frontend

import (
	"net/http"
	"sync/atomic"

	"github.com/caihua-yin/go-common/api"
	"github.com/gorilla/mux"
)

var (
	// Shutdown flags when HC should go down
	Shutdown int32
)

// HealthCheckHandler exposes store service health status as /_health
type HealthCheckHandler struct {
	api.Handler
}

// NewHealthCheckHandler creates new router for HealtCheckHandler
func NewHealthCheckHandler() http.Handler {
	result := &HealthCheckHandler{}
	result.Handler.Router = mux.NewRouter()

	result.Methods("GET").Path("/_health").Name("_health").HandlerFunc(result.HealthCheck)

	return result
}

// HealthCheck returns ok for now
func (h *HealthCheckHandler) HealthCheck(rw http.ResponseWriter, req *http.Request) {
	if atomic.LoadInt32(&Shutdown) > 0 {
		h.EmptyStatus(rw, 503)
		return
	}
}
