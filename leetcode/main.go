package main

func Stack(s string) bool {
	stack := make([]rune, 0)

	brackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, r := range s {
		if r == '(' || r == '[' || r == '{' {
			stack = append(stack, r)
		} else if len(stack) == 0 || stack[len(stack)-1] != brackets[r] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	//nums := []int{-5, 1, 5, 0, -7}

	println(Stack("([])0"))

	//fmt.Println(obj.SumRange(0, 3))
}

///
