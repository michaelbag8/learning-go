package main

import (
	"fmt"
	"math"
	"strings"
)

func printMultiplicationGrid(size int) {
	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}

type Student struct {
	Name  string
	Age   int
	Score int
}

type Book struct {
	Title  string
	Author string
	Pages  int
	Year   int
}

func printStudent(name string, age, score int) Student {
	return Student{
		Name:  name,
		Age:   age,
		Score: score,
	}
}

func countWords(word string) map[string]int {
	count := map[string]int{}
	w := strings.Fields(word)
	for _, words := range w {
		count[words]++
	}
	return count
}

func primeNumbers(max int) []int{
	var primes []int

	for i:= 2 ; i < max; i++{
		isPrime := true

		for j:= 2; j <= int(math.Sqrt(float64(i))); j++{
			if i % 2 == 0{
				isPrime = false
				break
			}
		}
		if isPrime{
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	fmt.Println("---------------Multiplication Table -------------")

	printMultiplicationGrid(12)

	fmt.Println("--------------------------------")
	person := Student{
		Name:  "Michael",
		Age:   100,
		Score: 98,
	}
	fmt.Println(person)
	fmt.Println("--------------------------------")

	person.Score = 99
	fmt.Println(person)

	book := Book{
		Title:  "The Notebook",
		Author: "John Greene",
		Pages:  267,
		Year:   2017,
	}
	fmt.Println(book)

	fmt.Println(printStudent("Michelle", 120, 100))

	fmt.Println("--------------------------")
	word := "the cat sat on the mat the cat"
	fmt.Println(countWords(word))
	fmt.Printf("unique word: %d\n", len(word))

	fmt.Println("-------Prime Numbers------")
	fmt.Println(primeNumbers(10))

}
