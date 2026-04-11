package main

import (
	"fmt"
	"strconv"
	"strings"
)

func makerPrint(word string) int {
	count := 1
	word = strings.ReplaceAll(word, ", ", ",")

	words := strings.Fields(word)
	
	for i := 0; i < len(words); i++ {

		if strings.HasPrefix(words[i], "(") && strings.HasSuffix(words[i], ")") {
			inner := strings.Trim(words[i], "()")
			if strings.Contains(inner, ",") {

				part := strings.Split(inner, ",")
				n, err := strconv.Atoi(strings.TrimSpace(part[1]))
				if err != nil || n < 1 {
					fmt.Println("Error", err)
				}
				count = n

			}

		}

	}

	return count
}

// maler functions that capture (marker, N)
func maker(word string) string {
	word = strings.ReplaceAll(word, ", ", ",")
	words := strings.Fields(word)

	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" {
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}
		if strings.HasPrefix(words[i], "(") && strings.HasSuffix(words[i], ")") {
			words[i] = strings.Trim(words[i], "()")
			if strings.Contains(words[i], ",") {
				inner := strings.Split(words[i], ",")
				n, err := strconv.Atoi(inner[1])
				if err == nil {
					start := i - n
					if start < 0 {
						start = 0
					}
					for j := start; j < i; j++ {
						words[j] = strings.ToUpper(words[j])
					}
					words = append(words[:i], words[i+1:]...)
					i--
				}
			}
		}

	}
	return strings.Join(words, " ")

}


func main() {
	word := "This is so exciting (up,4)"

	fmt.Println(makerPrint(word))

	fmt.Println(maker("This is so exciting (up)"))
	fmt.Println(maker("This is so exciting so big so small (up, 3)"))

}

