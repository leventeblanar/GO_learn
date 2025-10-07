package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Product struct {
	Name string
}

func (p Product) String() string { // hogy %s-sel is lehessen írni
	return p.Name
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var productList []Product

	fmt.Println("Enter products to create product list. Type 'done' to finish.")

	for {
		fmt.Print("> ") // kis prompt
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("read error:", err)
		}
		product := strings.TrimSpace(line)

		if strings.EqualFold(product, "done") {
			break
		}
		if product == "" {
			if len(productList) == 0 {
				fmt.Println("The list is empty.")
			} else {
				fmt.Println("Enter a product.")
			}
			continue
		}

		productList = append(productList, Product{Name: product})
	}

	if len(productList) == 0 {
		fmt.Println("The list is empty. Nothing to show.")
		return
	}

	fmt.Println("\nYour shopping list:")
	for i, p := range productList {
		fmt.Printf("%d. %s\n", i+1, p)
	}
	fmt.Println("Total items:", len(productList))
}
