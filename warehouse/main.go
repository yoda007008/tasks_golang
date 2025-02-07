package main

type Product struct { // структура товара на складе
	ID       int
	Name     string
	Quantity int
}

type Storage struct { // содержание склада
	Products map[int]Product
}

func main() {
}
