package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	assert.True(t, isPrime(7))
	assert.False(t, isPrime(8))
}

func isPrime(number int) bool {
	isPrime := true

	if number <= 1 {
		return false
	}

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return isPrime
}

func TestIsPalindrome(t *testing.T) {
	assert.True(t, isPalindrome("malam"))
	assert.False(t, isPalindrome("hari"))

}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if string(s[i]) != string(s[len(s)-i-1]) {
			return false
		}
	}

	return true
}

func TestFibonacci(t *testing.T) {
	assert.Equal(t, []int{0, 1, 1, 2, 3, 5, 8}, fibonacci(7))
}

func fibonacci(number int) []int {
	result := []int{}
	first := 0
	second := 1
	next := first + second

	for i := 0; i < number; i++ {
		result = append(result, first)
		next = first + second
		first = second
		second = next
	}

	return result
}

func TestFactorial(t *testing.T) {
	assert.Equal(t, 24, factorial(4))
}

func factorial(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result *= i
	}

	return result
}

func TestGetHighestNumber(t *testing.T) {
	assert.Equal(t, 200, getHighestNumber([]int{120, 180, 150, 125, 175, 200, 130}))
}

func getHighestNumber(numbers []int) int {
	sort.Ints(numbers)

	if len(numbers) > 0 {
		return numbers[len(numbers)-1]
	}

	return 0
}

func TestChangeStringInSlices(t *testing.T) {
	slices := []string{"dia", "suka", "membaca"}
	old := "membaca"
	new := "makan"

	changeStringInSlices(slices, old, new)

	assert.Equal(t, []string{"dia", "suka", "makan"}, slices)
}

func changeStringInSlices(slices []string, old, new string) {
	for k, v := range slices {
		if v == old {
			slices[k] = new
		}
	}
}
