package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("введите опрерацию из 2х строк в кавычках и через пробел или из строки и числа(если это `*` или `/`):")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') //ждет ввода данных
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
		os.Exit(1)
	}
	input = strings.TrimSpace(input) //очищает от пробелов
	fmt.Println(input)

	parts := strings.Split(input, "\" ")
	if len(parts) != 2 {
		fmt.Println("Неверный формат операции")
		os.Exit(1)
	}

	var k string
	var s int
	if parts[0][0:1] == "\"" && len(parts[0]) <= 11 && parts[1][0:1] == "+" && parts[1][2:3] == "\"" && len(parts[1]) <= 14 {
		k = parts[0] + parts[1][2:]
		k = strings.Replace(k, "\"", "", 4)
		fmt.Printf("%q", k)
	} else if parts[0][0:1] == "\"" && len(parts[0]) <= 11 && parts[1][0:1] == "/" && parts[1][2:3] != "\"" {
		parts[0] = strings.Replace(parts[0], "\"", "", -1)
		z, err := strconv.Atoi(parts[1][2:])
		if err == nil && z <= 10 {
			s = len(parts[0]) / z
			fmt.Printf("%q", parts[0][0:s])
		} else {
			fmt.Print("деление только на целое число <=10")
		}
	} else if parts[0][0:1] == "\"" && len(parts[0]) <= 11 && parts[1][0:1] == "*" && parts[1][2:3] != "\"" {
		parts[0] = strings.Replace(parts[0], "\"", "", -1)
		z, err := strconv.Atoi(parts[1][2:])
		if err == nil && z <= 10 {
			for i := 1; i <= z; i++ {
				k += parts[0]
			}
			if len(k) >= 42 {
				fmt.Print("\"" + k[0:40] + "..." + "\"")
			} else {
				fmt.Printf("%q", k)
			}
		} else {
			fmt.Print("умножение только на целое число <=10")
		}
	} else if parts[0][0:1] == "\"" && len(parts[0]) <= 11 && parts[1][0:1] == "-" && parts[1][2:3] == "\"" && len(parts[1]) <= 14 {
		parts[0] = strings.Replace(parts[0], "\"", "", -1)
		parts[1] = strings.Replace(parts[1], "\"", "", -1)
		if strings.Contains(parts[0], parts[1][2:]) == false {
			fmt.Printf("%q", parts[0])
		} else {
			k := strings.Index(parts[0], parts[1][2:])
			s = k + len(parts[1][2:])
			fmt.Print("\""+parts[0][0:k], parts[0][s:]+"\"")

		}
	} else {
		fmt.Print("ошибка ввода")

	}
}
