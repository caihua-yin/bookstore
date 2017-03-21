package model

import (
	"fmt"
	"github.com/satori/go.uuid"
	"testing"
)

func TestTimeConsuming(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	fmt.Println("This is a time consuming test")
}

func TestStoreItem(t *testing.T) {
	si := &StoreItem{
		ID:          uuid.NewV4().String(),
		Brand:       "Apple",
		Name:        "iPhone",
		Description: "The best phone in the planet",
	}
	if si.Name != "iPhone" {
		t.Error("Expected 'iPhone', got ", si.Name)
	}
}
