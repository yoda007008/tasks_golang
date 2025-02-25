package main

import "fmt"

func main() {
	list := make([]int, 4, 4)
	list = Append(list, 1)
	fmt.Println(list, len(list), cap(list))
}

func Append(list []int, elem int) []int { // реализация функции append
	var res []int

	resLen := len(list) + 1

	if resLen <= cap(list) { // проверка нужно ли увеличивать
		res = list[:resLen]
	} else {
		resCap := resLen
		if resCap < 3*len(list) { // увеличение емкости в 2 раза
			resCap = 3 * len(list)
		}
		res = make([]int, resLen, resCap) // создание нового массива res и копирование в него старый list
		copy(res, list)
	}
	res[len(list)] = elem // добавление нового элемента

	return res // возврат результата
}
