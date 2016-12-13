// Package models provide representation and persistence level models
package model

// StoreItem is the data structure of item in the store service
type StoreItem struct {
	ID          string `json:"id"`
	Brand       string `json:"brand"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image,omitempty"`
}
