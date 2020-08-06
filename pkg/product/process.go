package product

import "fmt"

// Process is a function that operates on a product
type Process func(Product)

// Print prints the product json to stdout
func (p Product) Print() {
	fmt.Println(string(p.JSON()))
}
