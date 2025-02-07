package main

import "fmt"

type Product struct { // структура товара на складе
	ID       int
	Name     string
	Quantity int
}

type Storage struct { // содержание склада
	Products map[int]Product
}

type Warehouse interface {
	AddProduct(product Product) error
	UpdateProduct(productID int, quantity int)
}

func (st *Storage) AddProduct(product Product) error {
	if st.Products == nil {
		st.Products = make(map[int]Product)
		fmt.Println("Продукт успешно добавлен")
	}
	if _, ok := st.Products[product.ID]; ok {
		return fmt.Errorf("Товар с ID %d уже существует на складе", product.ID)
	}
	st.Products[product.ID] = product
	return nil
}

func main() {
}
