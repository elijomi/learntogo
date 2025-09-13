package main

import "fmt"

/* Collection of small introductory exercises by ChatGPT do learn Go */
func main() {
	fmt.Println(doubleAll([]int{1, 2, 3})) // [2 4 6]
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
