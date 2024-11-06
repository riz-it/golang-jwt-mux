package productcontroller

import (
	"net/http"

	"github.com/riz-it/go-jwt-mux/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {

	products := []map[string]interface{}{
		{"id": 1, "name": "Product 1"},
		{"id": 2, "name": "Product 2"},
		{"id": 3, "name": "Product 3"},
	}

	helper.ResponseJSON(w, http.StatusOK, products)
}
