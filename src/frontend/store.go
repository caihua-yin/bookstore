package frontend

import (
	"net/http"

	"model"

	"github.com/caihua-yin/go-common/api"
	"github.com/caihua-yin/go-common/logging"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"github.com/uber-go/zap"
)

// StoreHandler handles job store API
type StoreHandler struct {
	api.Handler
	// User memory persistence here
	storeItems map[string]*model.StoreItem
}

// NewStoreHandler creates a StoreHandler
func NewStoreHandler() (*StoreHandler, error) {

	result := &StoreHandler{
		storeItems: make(map[string]*model.StoreItem),
	}
	result.Handler.Router = mux.NewRouter()

	// Post a new item
	result.Methods("POST").Path("/store/items").
		Name("post_item").
		HandlerFunc(result.PostItem)

	// Get an item info
	result.Methods("GET").Path("/store/items/{id}").
		Name("get_item").
		HandlerFunc(result.GetItem)

	return result, nil
}

// PostItem adds an new item to store
func (h *StoreHandler) PostItem(rw http.ResponseWriter, req *http.Request) {
	var bind struct {
		Brand       string `json:"brand"`
		Name        string `json:"name"`
		Description string `json:"description"`
		ImageURL    string `json:"image,omitempty"`
	}

	// Get POST request binds
	api.Check(api.Bind(req, &bind))
	item := &model.StoreItem{
		ID:          uuid.NewV4().String(),
		Brand:       bind.Brand,
		Name:        bind.Name,
		Description: bind.Description,
		ImageURL:    bind.ImageURL,
	}
	logger := logging.Logger()
	logger.Info("Add new item",
		zap.String("ID", item.ID),
		zap.String("Brand", item.Brand),
		zap.String("Name", item.Name),
		zap.String("Description", item.Description),
		zap.String("Image", item.ImageURL),
	)
	h.storeItems[item.ID] = item
	// Return 201 response with item ID in JSON
	h.JSONStatus(rw, 201, map[string]string{"id": item.ID})
}

// GetItem get an item by ID
func (h *StoreHandler) GetItem(rw http.ResponseWriter, req *http.Request) {
	var bind struct {
		ID string `mux:"id"`
	}
	api.Check(api.Bind(req, &bind))
	logger := logging.Logger()
	logger.Info("Get item", zap.String("ID", bind.ID))
	h.JSON(rw, h.storeItems[bind.ID])
}
