package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func intToRoman(num int) string {
	intNumbers := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	keySlice := make([]int, 0)
	for key, _ := range intNumbers {
		keySlice = append(keySlice, key)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keySlice)))
	fmt.Println(keySlice)
	result := ""
	for i := 0; i < len(keySlice); i++ {
		if num >= keySlice[i] {
			resInt := num / keySlice[i]
			num = num % keySlice[i]
			for j := 0; j < resInt; j++ {
				result += intNumbers[keySlice[i]]
			}
		}
	}
	return result
}

func romanToInt(s string) int {
	romanNumbers := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := romanNumbers[s[len(s)-1]]
	if len(s) == 1 {
		return result
	}
	for i := len(s) - 2; i >= 0; i-- {
		left := romanNumbers[s[i]]
		right := romanNumbers[s[i+1]]
		if left >= right {
			result += left
		} else {
			result -= left
		}
	}
	return result
}

func parseLineFromConsole() (int, byte, int, bool, bool, string) {
	/*
		берём строку,
		удаляем пробелы
		парсим (по пхпшьи) эксплоудом (по-питоньи сплитим)
		меняем тип у 0-го и 2-го
		и возвращаем три аргумента
	*/
	scanner := bufio.NewScanner(os.Stdin)
	text := ""
	slc := make([]string, 3)
	for {
		scanner.Scan()
		text = scanner.Text()
		if len(text) != 0 {
			slc = strings.Split(strings.TrimSpace(text), " ")
		} else {
			return 0, '+', 0, true, true, "input is empty"
		}
		if len(slc) != 3 {
			return 0, '+', 0, true, true, "input is not operation or got multiple operations"
		}
		break // в задании гарантирована одна строка, поэтому брейк на первой итерации
	}
	left, err := strconv.Atoi(slc[0])
	isRomanLeft := false
	isRomanRight := false
	if err != nil {
		left = romanToInt(slc[0])
		isRomanLeft = true
	}

	operator := make([]byte, 1)
	copy(operator, slc[1])

	right, err := strconv.Atoi(slc[2])
	if err != nil {
		right = romanToInt(slc[2])
		isRomanRight = true
	}
	if isRomanLeft == isRomanRight {
		if isRomanLeft && '-' == operator[0] && left <= right {
			return 0, '+', 0, true, true, "there is no negative or zero numbers in Roman numeric system"
		}
		return left, operator[0], right, isRomanLeft, false, "success"
	}
	return 0, '+', 0, true, true, "roman and arabic numbers got"
}

func evaluateOperation() (int, bool, error) {
	left, operator, right, isRoman, isError, msg := parseLineFromConsole()
	if isError {
		return 0, isRoman, errors.New(msg)
	}
	switch operator {
	case '+':
		return left + right, isRoman, nil
	case '-':
		return left - right, isRoman, nil
	case '*':
		return left * right, isRoman, nil
	case '/':
		return left / right, isRoman, nil
	}
	return 0, isRoman, errors.New("got other operation or operation is incorrect")
}

func main() {
	res, isRoman, err := evaluateOperation()
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		if isRoman {
			fmt.Printf("%s", intToRoman(res))
		} else {
			fmt.Printf("%d", res)
		}
	}
}
