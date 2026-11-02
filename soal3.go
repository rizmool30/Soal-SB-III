package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Name  string
	Price int
}

func main() {
	fmt.Println(getPurchasableProducts([]Product{
		{"milk", 9000},
		{"apple", 5000},
		{"cake", 15000},
		{"candybar", 2000},
		{"ice cream", 10000},
	}, 20000)) // [{candybar 2000},{apple 5000},{milk 9000}]

	fmt.Println(getPurchasableProducts([]Product{
		{"candy", 100},
		{"snack", 200},
		{"coffee", 300},
	}, 500)) // [{snack 200},{"coffee", 300}]

	fmt.Println(getPurchasableProducts([]Product{
		{"milk", 10000},
		{"mango", 20000},
	}, 5000)) // []

	fmt.Println(getPurchasableProducts([]Product{
		{"rice bowl", 1000},
		{"cupcake", 2000},
	}, 1000)) // [{rice bowl 1000}]

	fmt.Println(getPurchasableProducts([]Product{
		{"tea", 500},
		{"chai", 500},
		{"milk", 1000},
	}, 1000)) // [{tea 500},{chai 500}]

	fmt.Println(getPurchasableProducts([]Product{}, 10000)) // []
}

func getPurchasableProducts(products []Product, budget int) []Product {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})
	var hasil []Product
	total := 0

	for _, p := range products {
		if total+p.Price <= budget {
			hasil = append(hasil, p)
			total += p.Price
		} else {
			break
		}
	}

	return hasil

}
