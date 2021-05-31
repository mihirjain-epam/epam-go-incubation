package product

import (
	"encoding/json"
	"net/http"
	// "github.com/pluralsight/webservices/cors"
)

const productsPath = "products"

func handleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		products := getProductList()
		productsJson, err := json.Marshal(products)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(productsJson)
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodDelete:
	default:
	}

}

func handleProduct(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodDelete:
	default:
	}
}

// SetupRoutes :
func SetupRoutes(apiBasePath string) {

}
