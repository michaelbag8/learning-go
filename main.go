package main

import (
	"fmt"
	"strconv"
	"strings"
)

// replacing a word that start with vowel letter to "an"
func appendAn(str string) string {
	sd := []rune(str)
	first := sd[0]
	if len(sd) == 0 {
		return "a"
	}
	if first == 'a' || first == 'e' || first == 'i' || first == 'o' || first == 'u' || first == 'h' {
		return "an"
	}
	return "a"
}

// Puting "an" front of a word that start vowel letter
func fixArticle(s string) string {
	sd := strings.Fields(s)
	vowels := "aeiuohAEIOUH"
	for i := 0; i < len(sd)-1; i++ {
		if sd[i] == "a" && strings.ContainsRune(vowels, rune(sd[i+1][0])) {
			sd[i] = "an"
		}
	}
	return strings.Join(sd, " ")
}

// Checking for puntuation
func IsPuntuation(st string) bool {
	charset := ",.;:'?!"
	return strings.ContainsAny(st, charset)
}

// cleaning string
func removePuntuation(sd []string) string {
	var res strings.Builder
	for i, r := range sd {
		if i == 0 {
			res.WriteString(r)
			continue
		}
		if IsPuntuation(r) {
			res.WriteString(r)
		} else {
			res.WriteString(" ")
			res.WriteString(r)
		}
	}
	return res.String()
}

// Turning N word to uppercase
func LastN(c []string, n int) []string {
	for i := len(c) - n; i < len(c); i++ {
		if n > len(c) {
			n = len(c)
		}
		c[i] = strings.ToUpper(c[i])

	}
	return c

}

// Capitalizing words
func capitalizeWord(words string) string {
	return strings.ToUpper(words[:1]) + strings.ToLower(words[1:])
}

// Binary to Base 10
func binToDec(str string) (int64, error) {
	con, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		return 0, err
	}
	return con, nil
}

// Hexedecimal to base 10
func hexToDec(str string) (int64, error) {
	con, err := strconv.ParseInt(str, 16, 64)
	if err != nil {
		return 0, err
	}
	return con, nil
}

// Calling all the functions
func main() {
	ds := "Welcome Mr a apple, you came with Mr a orange at a hour that is not suitable for Mr a umbralla"
	cs := []string{"yo", "ye", "ya", "yu", "yi"}
	fmt.Println(LastN(cs, 3))
	fmt.Println(fixArticle(ds))

	fmt.Println(appendAn("apple"))
	fmt.Println(appendAn("hour"))

	fmt.Println(removePuntuation([]string{"hello", ",", "world"}))

	fmt.Println(capitalizeWord("WORD"))
	fmt.Println(capitalizeWord("hello"))

	fmt.Println(binToDec("1111111"))
	fmt.Println(binToDec("1010"))

	fmt.Println(hexToDec("FF"))
	fmt.Println(hexToDec("1E"))
}
