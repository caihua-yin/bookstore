package frontend

import (
	"net/http"

	"github.com/caihua-yin/go-common/api"
	"github.com/gorilla/mux"
)

// VersionHandler exposes store service version as /_version
type VersionHandler struct {
	api.Handler
}

// NewVersionHandler creates new router for VersionHandler
func NewVersionHandler() http.Handler {
	result := &VersionHandler{}
	result.Handler.Router = mux.NewRouter()

	result.Methods("GET").Path("/_version").HandlerFunc(result.Version)

	return result
}

// Version returns current version
func (h *VersionHandler) Version(rw http.ResponseWriter, req *http.Request) {
	h.JSON(rw, struct {
		Version string `json:"version"`
	}{
		Version: Version,
	})
}
