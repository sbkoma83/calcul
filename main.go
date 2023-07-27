package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("введите мат.опрерацию из двух арабских или римских чисел через пробел:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') //ждет ввода данных
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
		os.Exit(1)
	}
	input = strings.TrimSpace(input) //очищает от пробелов
	fmt.Println(input)

	// Разбиваем строку на операнды и оператор
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		fmt.Println("Неверный формат операции")
		os.Exit(1)
	}
	var rome = [101]string{"ошибка в римских числах нет 0 и отрицательных чисел", "I", "II", "III", "IV",
		"V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
		"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIV", "XXX", "XXXI", "XXXII", "XXXIII",
		"XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL", "XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI",
		"XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX", "LXI",
		"LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX", "LXXI", "LXXII", "LXXIII", "LXXIV",
		"LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX", "LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV",
		"LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC", "XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII",
		"XCVIII", "XCIX", "C"}
	var a, b, k int

	z, err := strconv.Atoi(parts[0])
	if err != nil {
		z = a
	}

	x, err := strconv.Atoi(parts[2])
	if err != nil {
		x = b
	}

	if err == nil && z != a {

		if z >= 0 && z <= 10 && parts[1] == "+" && x >= 0 && x <= 10 {
			fmt.Print(z + x)
		} else if z >= 0 && z <= 10 && parts[1] == "-" && x >= 0 && x <= 10 {
			fmt.Print(z - x)
		} else if z >= 0 && z <= 10 && parts[1] == "*" && x >= 0 && x <= 10 {
			fmt.Print(z * x)
		} else if z >= 0 && z <= 10 && parts[1] == "/" && x >= 0 && x <= 10 {
			fmt.Print(z / x)
		} else {
			fmt.Print("Ошибка ввода")
			os.Exit(1)
		}
	} else if z == a && x != b || z != a && x == b {
		fmt.Print("Не используйте римские и арабские числа вместе")
		os.Exit(1)
	} else {

		switch parts[0] {
		case "I":
			a = 1
		case "II":
			a = 2
		case "III":
			a = 3
		case "IV":
			a = 4
		case "V":
			a = 5
		case "VI":
			a = 6
		case "VII":
			a = 7
		case "VIII":
			a = 8
		case "IX":
			a = 9
		case "X":
			a = 10
		default:
			fmt.Println("Ошибка ввода")
			os.Exit(1)
		}
		switch parts[2] {
		case "I":
			b = 1
		case "II":
			b = 2
		case "III":
			b = 3
		case "IV":
			b = 4
		case "V":
			b = 5
		case "VI":
			b = 6
		case "VII":
			b = 7
		case "VIII":
			b = 8
		case "IX":
			b = 9
		case "X":
			b = 10
		default:
			fmt.Println("Ошибка ввода")
			os.Exit(1)
		}

		if parts[1] == "+" {
			k = a + b
			fmt.Print(rome[k])
		} else if parts[1] == "-" && a >= b {
			k = a - b
			fmt.Print(rome[k])
		} else if parts[1] == "-" && a < b {
			fmt.Print(rome[0])
		} else if parts[1] == "*" {
			k = a * b
			fmt.Print(rome[k])
		} else if parts[1] == "/" {
			k = a / b
			fmt.Print(rome[k])
		} else {
			fmt.Print("Ошибка ввода")
			os.Exit(1)
		}
	}
}
