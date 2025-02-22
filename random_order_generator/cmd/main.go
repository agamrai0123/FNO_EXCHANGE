package main

import (
	"fmt"

	"github.com/agamrai0123/FNO_EXCHANGE/random_order_generator/internal"
)

// Usage example:
func main() {
	// Generate a single random order
	// order := internal.GenerateRandomOrder()
	// fmt.Printf("Single models.Order: %+v\n", order)

	// Generate multiple random orders
	orders := internal.GenerateRandomOrders(5)
	fmt.Printf("Generated %d orders\n", len(orders))
}
