package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// (up)
func upperLastWord(word []string) []string {
	for i := 0; i < len(word); i++ {
		if word[i] == "(up)" {
			if i > 0 {
				word[i-1] = strings.ToUpper(word[i-1])
			}
			word = append(word[:i], word[i+1:]...)

		}

	}
	return word
}

// (up,3)
func toUpperLastWord(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" {
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}

		if strings.HasPrefix(words[i], "(up,") && strings.HasSuffix(words[i], ")") {
			numStr := strings.TrimSuffix(strings.TrimPrefix(words[i], "(up,"), ")")
			n, err := strconv.Atoi(numStr)
			if err == nil {
				start := i - n
				if start < 0 {
					start = 0
				}

				for j := start; j < i; j++ {
					words[j] = strings.ToUpper(words[j])
				}
			}

			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return words
}

// fix article
func fixArticle(word string) string {
	words := strings.Fields(word)
	vowels := "aeiouhAEIOUH"
	for i := 0; i < len(words)-1; i++ {
		if (words[i] == "a" || words[i] == "A") && strings.ContainsRune(vowels, rune(words[i+1][0])) {
			if words[i] == "A" {
				words[i] = "An"
			} else {
				words[i] = "an"
			}
		}

	}
	return strings.Join(words, " ")

}

// checking for puntuation
func isPunctuation(word string) bool {
	charset := ";,.!:"
	return strings.ContainsAny(word, charset)
}

// removing puntuation
func removePuntuation(word []string) string {
	var result strings.Builder
	for i, ch := range word {
		if i == 0 {
			result.WriteString(ch)
		}
		if isPunctuation(ch) {
			result.WriteString(ch)
		} else {
			result.WriteString(" ")
			result.WriteString(ch)
		}
	}
	return result.String()
}

// upper case the base on n
func upperLastN(word []string, n int) []string {
	for i := len(word) - n; i < len(word); i++ {
		if n < 0 {
			n = len(word)
		}
		word[i] = strings.ToUpper(word[i])
	}
	return word

}

// convert to base
func ConvertNumbers(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			n, err := strconv.ParseInt(words[i-1], 16, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(n, 10)
			}
			words = append(words[:i], words[i+1:]...)
			i--
			continue

		}
		if words[i] == "(bin)" && i > 0 {
			n, err := strconv.ParseInt(words[i-1], 2, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(n, 10)
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}

	}

	return words

}

func capitalize(word string) string {
	words := strings.Fields(word)
	for i, ch := range words {
		runes := []rune(strings.ToLower(ch))
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}

		words[i] = string(runes)
	}

	return strings.Join(words, " ")
}

func main() {
	num := []string{"10", "(bin)", "years"}
	num2 := []string{"1E", "(hex)", "files"}

	fmt.Println(ConvertNumbers(num))
	fmt.Println(ConvertNumbers(num2))
	fmt.Println("----------------------------")

	w := []string{"hello", "come", "down", "(up,3)"}
	fmt.Println(toUpperLastWord(w))
	fmt.Println(isPunctuation(","))

	fmt.Println(removePuntuation([]string{"hello", ",", "world"}))
	fmt.Println("----------------------------")

	fmt.Println(fixArticle("this is a orange in the a hour of need"))
	fmt.Println("----------------------------")

	word := []string{"hello", "come", "down", "(up)"}
	wordz := []string{"hello", "(up)", "down"}
	wordx := []string{"hello", "rise", "down", "got", "wanted", "okay"}

	fmt.Println(upperLastN(wordx, 3))
	fmt.Println(toUpperLastWord(word))
	fmt.Println(toUpperLastWord(wordz))
}
