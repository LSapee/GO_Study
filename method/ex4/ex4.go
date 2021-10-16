package main

import "fmt"

type Cart struct {
	productList string
}

func (c *Cart) AddProduct(product string) {
	if c.productList != "" {
		c.productList += ", "
	}
	c.productList += product
}

func (c *Cart) Clear() {
	c.productList = ""
}

func (c Cart) GetProductList() string {
	return c.productList
}

func main() {
	c := &Cart{}
	c.AddProduct("apple")
	c.AddProduct("kimchi")

	fmt.Println(c.GetProductList())

	c.Clear()
	c.AddProduct("watermelon")
	fmt.Println(c.GetProductList())

	//결과 apple, kimchi
	//결과 2 watermelon
}
