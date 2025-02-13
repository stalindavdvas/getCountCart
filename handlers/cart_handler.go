package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// GetCartCount returns the total number of items in the cart
func GetCartCount(client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := "user:1" // Clave única para el carrito del usuario

		// Obtener todos los productos del carrito
		cartItems, err := client.HGetAll(ctx, userID).Result()
		if err != nil {
			http.Error(w, "Error al obtener los datos del carrito", http.StatusInternalServerError)
			return
		}

		// Calcular la cantidad total de productos
		total := 0
		for _, productJSON := range cartItems {
			// Decodificar el JSON almacenado en Redis
			var productData map[string]interface{}
			err := json.Unmarshal([]byte(productJSON), &productData)
			if err != nil {
				http.Error(w, "Error al procesar los datos del carrito", http.StatusInternalServerError)
				return
			}

			// Extraer la cantidad del producto
			quantity, ok := productData["quantity"].(float64) // Redis devuelve números como float64
			if !ok {
				http.Error(w, "Error al leer la cantidad del producto", http.StatusInternalServerError)
				return
			}

			// Sumar la cantidad al total
			total += int(quantity)
		}

		// Preparar la respuesta
		response := map[string]int{
			"count": total,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
