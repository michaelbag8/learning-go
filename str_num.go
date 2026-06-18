package main

import (
	"errors"
	"fmt"
	"unicode"

	"golang.org/x/telemetry/counter"
)

var ErrDivisibleByZero = errors.New("Not divisible by zero")

func reverseNum(num int) int {
	reversed := 0
	for num != 0 {
		digit := num % 10
		reversed = reversed*10 + digit
		num /= 10
	}
	return reversed
}

func CountEvenNum(num []int) int {
	CountEven := 0
	for _, n := range num {
		if n%2 == 0 {
			CountEven++
		}
	}
	return CountEven
}

func sumEvenNums(num []int) int {
	sum := 0
	for _, n := range num {
		if n%2 == 0 {
			sum += n
		}
	}
	return sum
}

func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, ErrDivisibleByZero
	}
	return float64(a / b), nil
}

func extractNum(str string) string {
	runes := []rune(str)

	for i, ch := range runes {
		if unicode.IsDigit(ch) {
			runes[i] = 'X'
		}

	}
	return string(runes)
}

func compressChar(str string) string {
	var result strings.Builder
	counter := 1
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			counter++
		} else {
			result.WriteByte(str[i-1])
			if counter > 1 {
				result.WriteString(strconv.Itoa(counter))
			}
			counter = 1

		}

	}
	result.WriteByte(str[len(str)-1])
	if counter > 1 {
		result.WriteString(strconv.Itoa(counter))
	}
	return result.String()
}

func main() {
	fmt.Println(reverseNum(12345))

	fmt.Println(Divide(5, 0))

	fmt.Println(CountEvenNum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))

	fmt.Println(sumEvenNums([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))

	fmt.Println(extractNum("ahnd45mnsd3ased9"))

	fmt.Println(compressChar("aaaaabcccccccccdeeeee"))

}
