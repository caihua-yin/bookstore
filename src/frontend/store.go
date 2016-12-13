package frontend

import (
	"log"
	"net/http"

	"model"

	"github.com/caihua-yin/go-common/api"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
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
	log.Printf("Add new item, ID: %s, brand: %s, name: %s, description: %s, image: %s",
		item.ID, item.Brand, item.Name, item.Description, item.ImageURL)
	h.storeItems[item.ID] = item
	// Return 201 response with item ID in JSON
	h.JSONStatus(rw, 201, map[string]string{"ID": item.ID})
}

// GetItem get an item by ID
func (h *StoreHandler) GetItem(rw http.ResponseWriter, req *http.Request) {
	var bind struct {
		ID string `mux:"id"`
	}
	api.Check(api.Bind(req, &bind))
	h.JSON(rw, h.storeItems[bind.ID])
}
