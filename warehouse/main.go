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

func (st *Storage) UpdateProduct(productID int, quantity int) error {
	product, ok := st.Products[productID]
	if !ok {
		return fmt.Errorf("Товар с ID %d не найден на складе", productID)
	}
	if product.Quantity+quantity < 0 {
		return fmt.Errorf("Количество товаров не может быть отрицательным")
	}
	product.Quantity += quantity
	st.Products[productID] = product
	return nil
}

func main() {
	storage := &Storage{
		Products: make(map[int]Product),
	}
	// добавление продукта
	err := storage.AddProduct(Product{ID: 1, Name: "Ноутбук", Quantity: 3})

	if err != nil {
		fmt.Println("Ошибка при добавлении товара", err)
		return
	}
	// изменение продукта
	err = storage.UpdateProduct(1, -5)
	if err != nil {
		fmt.Println("Ошибка в изменении товара", err)
	}
}
