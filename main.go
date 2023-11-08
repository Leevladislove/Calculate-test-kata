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

func checkCountArgs(args []string) error {
	if len(args) != 3 {
		return fmt.Errorf("Ошибка: Неверное количество аргументов")
	}
	return nil
}

func isNum(val string) bool {
	_, err := strconv.Atoi(val)
	return err == nil
}

func isRoman(val []string) bool {
	return checkRoman(val[0]) && checkRoman(val[2])
}

func checkRoman(val string) bool {
	romanNums := map[string]bool{
		"I": true, "II": true, "III": true, "IV": true, "V": true,
		"VI": true, "VII": true, "VIII": true, "IX": true, "X": true,
	}
	return romanNums[val]
}

// конверт взять из интернета
func romanToArabic(roman string) (int, error) {
	romanNums := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}

	if val, ok := romanNums[roman]; ok {
		return val, nil
	}

	return 0, fmt.Errorf("Ошибка: Введите римское число от I до X")
}

// конверт взять из интернета
func arabicToRoman(num int) (string, error) {
	if num < 1 {
		return "", fmt.Errorf("Ошибка: Римское число не может быть 0 или отрицательным")
	}

	romanNums := map[int]string{
		1000: "M", 900: "CM", 500: "D", 400: "CD", 100: "C",
		90: "XC", 50: "L", 40: "XL", 10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I",
	}

	var result strings.Builder

	for value, numeral := range romanNums {
		for num >= value {
			num -= value
			result.WriteString(numeral)
		}
	}

	return result.String(), nil
}

func checkRangeNums(num1, num2 int) error {
	if (num1 < 1 || num1 > 10) || (num2 < 1 || num2 > 10) {
		return fmt.Errorf("Ошибка: Число должно быть в диапазоне от 1 до 10")
	}
	return nil
}

func calculateArabAndRoman(exp []string) (interface{}, error) {
	if err := checkCountArgs(exp); err != nil {
		return nil, err
	}

	var num1, num2 int
	var err1, err2 error

	if isNum(exp[0]) && isNum(exp[2]) {
		num1, err1 = strconv.Atoi(exp[0])
		num2, err2 = strconv.Atoi(exp[2])
	} else if isRoman(exp) {
		num1, err1 = romanToArabic(exp[0])
		num2, err2 = romanToArabic(exp[2])
	} else {
		return nil, fmt.Errorf("Ошибка: Неверный формат чисел")
	}

	if err1 != nil || err2 != nil {
		return nil, fmt.Errorf("Ошибка: Введите целое число")
	}

	if err := checkRangeNums(num1, num2); err != nil {
		return nil, err
	}

	operator := exp[1]
	switch operator {
	case "+":
		result := num1 + num2
		if isRoman(exp) {
			romanResult, err := arabicToRoman(result)
			if err != nil {
				return "", err
			}
			return romanResult, nil
		}
		return result, nil
	case "-":
		result := num1 - num2
		if isRoman(exp) {
			romanResult, err := arabicToRoman(result)
			if err != nil {
				return "", err
			}
			return romanResult, nil
		}
		return result, nil
	case "*":
		result := num1 * num2
		if isRoman(exp) {
			romanResult, err := arabicToRoman(result)
			if err != nil {
				return "", err
			}
			return romanResult, nil
		}
		return result, nil
	case "/":
		result := num1 / num2
		if isRoman(exp) {
			romanResult, err := arabicToRoman(result)
			if err != nil {
				return "", err
			}
			return romanResult, nil
		}
		return result, nil
	default:
		return nil, fmt.Errorf("Ошибка: Неверный математический оператор")
	}
}

func main() {
	fmt.Println("Введите выражение: ")
	input := getInput()
	result, err := calculateArabAndRoman(input)

	if err == nil {
		fmt.Println("Результат:", result)
	} else {
		fmt.Println(err)
	}
}
