package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() []string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.Fields(strings.TrimSpace(input))
}

func checkRange(num1, num2 int) bool {
	fmt.Println("Число должно быть в диапазоне от 1 до 10")
	return num1 >= 1 && num1 <= 10 && num2 >= 1 && num2 <= 10
}

func checkCountArgument(arg []string) bool {
	if len(arg) != 3 {
		fmt.Println("Неверное количество аргументов")
		return false
	}
	return true
}

func calculate(expr []string) {
	if !checkCountArgument(expr) {
		return
	}

	num1, err1 := strconv.Atoi(expr[0])
	num2, err2 := strconv.Atoi(expr[2])
	
	if err1 != nil || err2 != nil {
		fmt.Println("Неверные операнды")
		return
	}

	if !checkRange(num1, num2) {
		return
	}

	operator := expr[1]
	switch operator {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Делить на 0 нельзя")
			return
		}
		fmt.Println(num1 / num2)
	default:
		fmt.Println("Неверный оператор")
	}
}

func main() {
	for {
		fmt.Println("Введите выражение: ")
		input := getInput()
		calculate(input)
	}
}
