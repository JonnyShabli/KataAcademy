package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите выражение: ")
	expression, _ := reader.ReadString('\n')

	operator, operand1, operand2, romeFlag := parse(expression)

	result := calculate(operator, operand1, operand2)
	if romeFlag == true {
		if result < 1 {
			fmt.Println("Для римских чисел допустимый результат должен быть больше 1")
			os.Exit(1)
		}
		fmt.Printf("Результат вычислений: %v\n", intToRoman(result))
	} else {
		fmt.Printf("Результат вычислений: %v\n", result)
	}

}

func parse(expression string) (string, int, int, bool) {
	//  удаляем символ перехода на новую строку
	//expression = strings.TrimSuffix(expression, "\n")
	expression = strings.TrimRight(expression, " \n")
	//	разделяем сроку на значения и оператор, проверяем корректность выражения
	tokens := strings.Split(expression, " ")
	if len(tokens) != 3 {
		fmt.Println("Выражение должно быть вида: два операнда и один оператор (+, -, /, *), введенные через пробел \n")
		os.Exit(1)
	}
	//	карта допустимых арабских чисел
	arab := map[string]bool{
		"1":  true,
		"2":  true,
		"3":  true,
		"4":  true,
		"5":  true,
		"6":  true,
		"7":  true,
		"8":  true,
		"9":  true,
		"10": true,
	}
	//	карта допустимых римских чисел
	rome := map[string]string{
		"I":    "1",
		"II":   "2",
		"III":  "3",
		"IV":   "4",
		"V":    "5",
		"VI":   "6",
		"VII":  "7",
		"VIII": "8",
		"IX":   "9",
		"X":    "10",
	}

	//	маркеры принадлежности значений к допустимым римским числам
	_, chk1 := rome[tokens[0]]
	_, chk2 := rome[tokens[2]]

	//	проверяем первое и второе значения
	if arab[tokens[0]] == false && chk1 == false {
		fmt.Println("Неверное значение 1: Допустимо использовать римские или арабские числа от 1 до 10 включительно")
		if arab[tokens[2]] == false && chk2 == false {
			fmt.Println("Неверное значение 2: Допустимо использовать римские или арабские числа от 1 до 10 включительно")
			os.Exit(1)
		}
		os.Exit(1)
	}

	//	проверяем второе значение
	if arab[tokens[2]] == false && chk2 == false {
		fmt.Println("Неверное значение 2: Допустимо использовать числа от 1 до 10 включительно")
		os.Exit(1)
	}

	if (chk1 == true && arab[tokens[2]] == true) || (arab[tokens[0]] == true && chk2 == true) {
		fmt.Println("Допустимо использовать либо римские либо арабские числа одновременно")
		os.Exit(1)
	}

	var operand1 int
	var operand2 int
	var romeFlag bool
	operator := tokens[1]

	//	разные сценарии для римских и арабских чисел
	switch {
	case arab[tokens[0]] == true && arab[tokens[2]]: // оба числа арабские
		operand1, _ = strconv.Atoi(tokens[0])
		operand2, _ = strconv.Atoi(tokens[2])
	case chk1 == true && chk2 == true: // оба числа римские
		operand1, _ = strconv.Atoi(rome[tokens[0]])
		operand2, _ = strconv.Atoi(rome[tokens[2]])
		romeFlag = true
	}
	return operator, operand1, operand2, romeFlag
}

// перевод арабских чисел в римские
func intToRoman(num int) string {
	r := [][]string{
		{"", "M", "MM", "MMM"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}}
	n := []int{1000, 100, 10, 1}
	result := ""
	for k, v := range n {
		result += r[k][num/v]
		num = num % v
	}
	return result
}

func calculate(operator string, operand1 int, operand2 int) int {
	var result int
	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		result = operand1 / operand2
	default:
		fmt.Println("Неверный оператор:", operator)
		os.Exit(1)
	}

	return result
}
