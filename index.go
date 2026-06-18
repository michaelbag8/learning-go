package main

import (
	"fmt"
	"strings"
	"unicode"
)

// vowels count
func vowelCount(word string) int {
	count := 0
	words := strings.ToLower(word)
	for _, ch := range words {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			count++
		}

	}

	return count
}

// words count
func wordsCount(word string) int {
	count := 0
	for i := 0; i < len(word); i++ {
		count++
	}
	return count
}

// words count
func WordCount(word string) int {
	words := strings.Fields(word)
	return len(words)
}

// uppercase
func upperCase(word string) string {
	return strings.ToUpper(word)
}

// capitalize each word
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

// check for palindrome
func checkPalindrone(word string) string {
	if word == reverse(word) {
		return "It is a palindrome"
	}
	return "Not a palindrome"
}

// reverse a string
func reverse(word string) string {
	runes := []rune(word)
	for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
		runes[l], runes[r] = runes[r], runes[l]
	}
	return string(runes)
}

// textanlyzer
func analyzeText(word string) (int, int, string) {
	countVowel := 0

	words := strings.Fields(word)

	countWords := len(words)

	wordz := strings.ToLower(word)
	for _, ch := range wordz {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			countVowel++
		}
	}

	for i, char := range words {
		runes := []rune(strings.ToLower(char))
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}
		words[i] = string(runes)
	}

	return countVowel, countWords, strings.Join(words, " ")
}

func main() {
	fmt.Println("Vowels count: ", vowelCount("Michael"))
	fmt.Println("Vowels count: ", vowelCount("welcome home my boy"))
	fmt.Println("Words count: ", wordsCount("Michael Bag"))
	fmt.Println("Another Words Count: ", WordCount("Thanks be to you"))
	fmt.Println(upperCase("hello my boy"))
	fmt.Println(checkPalindrone("michael"))
	fmt.Println(analyzeText("hello boy WELCOME BACK TO CAMPUS"))
}
