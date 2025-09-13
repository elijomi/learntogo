package main

import (
	"fmt"
	"learntogo/person"
	"strings"
)

/* Collection of small introductory exercises by ChatGPT do learn Go */
func main() {
	p1 := person.NewPerson("Alice", -2)
	p2 := person.NewPerson("Bob", 17)
	p3 := person.NewPerson("Charlie", 19)

	p1.AddFriend(p1)
	p1.AddFriend(p2)
	p1.AddFriend(p2)
	p1.AddFriend(p3)
	fmt.Println(p1.Greet())
	fmt.Println("Alice's friends before:", strings.Join(p1.FriendNames(), ", "))
}

// 5. Write a function countLetters(s string) map[rune]int that returns a map with the frequency of each letter in the string.
func countLetters(s string) map[rune]int {
	res := map[rune]int{}
	for _, r := range s {
		res[r]++
	}

	return res
}

// 4. Write a function doubleAll(nums []int) []int that takes a slice of ints and returns a new slice where every number is doubled.
func doubleAll(nums []int) []int {
	res := make([]int, len(nums)) // new slice same length as nums
	for i := 0; i < len(nums); i++ {
		res[i] = nums[i] * 2
	}

	return res
}

// 3. Write a function sumToN(n int) int that returns the sum of all numbers from 1 to n using a for loop.
func sumToN(n int) int {
	sum, i := 0, 0
	for i < n {
		i++
		sum += i
	}

	return sum
}

// 2: Write a function add(a int, b int) int that returns the sum.
func add(a int, b int) int {
	return a + b
}

// 1: Define a function hello(name string) string that returns "Hello, <name>!".
func hello(name string) string {
	return "Hello, " + name + "!"
}
