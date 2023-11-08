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

func isRoman(val string) bool {
	romanNumerals := map[string]bool{
		"I": true, "II": true, "III": true, "IV": true, "V": true,
		"VI": true, "VII": true, "VIII": true, "IX": true, "X": true,
	}
	return romanNumerals[val]
}

func isRomanNum(val []string) bool {
	return isRoman(val[0]) && isRoman(val[2])
}

func romanToArabic(roman string) (int, error) {
	romanNums := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}

	if val, ok := romanNums[roman]; ok {
		return val, nil
	}

	return 0, fmt.Errorf("Ошибка: Римское число должно быть от I до X")
}

func arabicToRoman(num int) (string, error) {
	romanNums := map[int]string{
		1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
		6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
	}

	if num < 1 || num > 10 {
		return "", fmt.Errorf("Ошибка: Результат римского числа должен быть от 1 до 10")
	}

	if numeral, ok := romanNums[num]; ok {
		return numeral, nil
	}

	return "", fmt.Errorf("Ошибка: Результат римского числа должен быть от 1 до 10")
}

func checkRangeNums(num1, num2 int) error {
	if (num1 < 1 || num1 > 10) || (num2 < 1 || num2 > 10) {
		return fmt.Errorf("Ошибка: Число должно быть в диапазоне от 1 до 10")
	}
	return nil
}

func calculateArabAndRoman(expr []string) (interface{}, error) {
	if err := checkCountArgs(expr); err != nil {
		return nil, err
	}

	var num1, num2 int
	var err1, err2 error

	if isNum(expr[0]) && isNum(expr[2]) {
		num1, err1 = strconv.Atoi(expr[0])
		num2, err2 = strconv.Atoi(expr[2])
	} else if isRomanNum(expr) {
		num1, err1 = romanToArabic(expr[0])
		num2, err2 = romanToArabic(expr[2])
	} else {
		return nil, fmt.Errorf("Ошибка: Неверный формат чисел")
	}

	if err1 != nil || err2 != nil {
		return nil, fmt.Errorf("Ошибка: Введите целое число")
	}

	if err := checkRangeNums(num1, num2); err != nil {
		return nil, err
	}

	operator := expr[1]
	switch operator {
	case "+":
		result := num1 + num2
		if isRomanNum(expr) {
			romanResult, err := arabicToRoman(result)
			if err != nil {
				return nil, err
			}
			return romanResult, nil
		}
		return result, nil
	case "-":
		result := num1 - num2
		if isRomanNum(expr) {
			romanResult, err := arabicToRoman(result)
			if err != nil {
				return nil, err
			}
			return romanResult, nil
		}
		return result, nil
	case "*":
		result := num1 * num2
		if isRomanNum(expr) {
			romanResult, err := arabicToRoman(result)
			if err != nil {
				return nil, err
			}
			return romanResult, nil
		}
		return result, nil
	case "/":
		result := num1 / num2
		if isRomanNum(expr) {
			romanResult, err := arabicToRoman(result)
			if err != nil {
				return nil, err
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
