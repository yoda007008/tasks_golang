package main

import "fmt"

func main() {
	x := -5.0
	fmt.Println("     x  |  y")
	for x <= 5 {
		y := 5 - x*x/2
		fmt.Printf("%6.2f  | %6.2f\n", x, y) // вот эту строчку если честно я сформатировал при помощи нейронки, а так писал сам
		x = x + 0.5
	}

}
