package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// GetCartCount returns the total number of items in the cart
func GetCartCount(client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := "user:1" // Unique key for the user's cart

		// Get all items in the cart
		cartItems, err := client.HGetAll(ctx, userID).Result()
		if err != nil {
			http.Error(w, "Error fetching cart data", http.StatusInternalServerError)
			return
		}

		// Calculate the total quantity
		total := 0
		for _, quantity := range cartItems {
			qty, _ := strconv.Atoi(quantity)
			total += qty
		}

		// Prepare the response
		response := map[string]int{
			"count": total,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
