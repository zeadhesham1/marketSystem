package main

import (
	"fmt"
)

type Product struct {
	id       int
	name     string
	price    float32
	quantity int
}

//
func printProduct(aproduct Product) {
	fmt.Printf("%v      %v    %v    %v\n", aproduct.id, aproduct.name, aproduct.price, aproduct.quantity)
}

func printAllProducts(productArr []Product) {

	for p := range productArr {
		printProduct(productArr[p])
	}
}

func main() {
	aproduct1 := Product{2, "cheese", 3.4, 5}
	aproduct6 := Product{7, "Mushroom", 8.9, 10}
	aproduct11 := Product{12, "milk", 13.14, 15}
	aproduct16 := Product{17, "water", 18.19, 20}
	productArray := []Product{aproduct1, aproduct6, aproduct11, aproduct16}

	menuChoice := chooseFromMenu()

	switch menuChoice {
	case 1:
		printAllProducts(productArray)
	case 2:
		fmt.Print("3. find existing product\n")

	case 3:
		fmt.Print("3. find existing product\n")
	case 4:
		fmt.Print("4. edit existing product\n")
	case 5:
		fmt.Print("5. delete existing product\n ")
	case 6:
		fmt.Print("6. exit\n")
	}

}
func chooseFromMenu() int {
	var input int
	fmt.Print("choose one of the following operations:\n")
	fmt.Print("1. view all products\n")
	fmt.Print("2. add new product\n")
	fmt.Print("3. find existing product\n")
	fmt.Print("4. edit existing product\n")
	fmt.Print("5. delete existing product\n ")
	fmt.Print("6. exit\n")
	fmt.Print("enter your choice :")
	// int validation required
	fmt.Scanln(&input)
	return input
}
