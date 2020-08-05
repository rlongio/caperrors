package product

import "fmt"

// Process is a function that operates on a product
type Process func(Product)

// Print print the product to stdout
func (p Product) Print() {
	j, err := p.JSON()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(j))
}
