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
func main() {
	word := "This is so exciting (up,4)"

	fmt.Println(makerPrint(word))

}

