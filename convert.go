package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hexToDecimal(s string) string {
	data, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return s
	}
	return strconv.Itoa(int(data))

}

func binToDecimal(s string) string {
	data, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return s
	}
	return strconv.Itoa(int(data))

}

func toDecimal(word string, base int) string {
	original := word
	word = strings.TrimSpace(word)

	negative := false
	if strings.HasPrefix(word, "-") {
		negative = true
		word = strings.TrimPrefix(word, "-")
	}

	data, err := strconv.ParseInt(word, base, 64)
	if err != nil {
		return original
	}
	if negative {
		data = -data
	}

	return strconv.Itoa(int(data))

}

func main() {
	fmt.Println(hexToDecimal("1E"))
	fmt.Println(hexToDecimal("ff"))
	fmt.Println(hexToDecimal("xyz"))

	fmt.Println(binToDecimal("10"))
	fmt.Println(binToDecimal("1101"))
	fmt.Println(binToDecimal("999"))

	fmt.Println(toDecimal("10", 2))
	fmt.Println(toDecimal("0644", 8))
	fmt.Println(toDecimal("  1E  ", 16))
	fmt.Println(toDecimal("-ff", 16))

}
