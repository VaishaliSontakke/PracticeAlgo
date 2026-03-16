package main

import "fmt"

func main() {
	fmt.Println(isValid("(()"))
}

func isValid(s string) bool {
	stack := []rune{}
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (top == '(' && ch != ')') || (top == '[' && ch != ']') || (top == '{' && ch != '}') {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
