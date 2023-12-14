package main

import (
	"fmt"
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
	assert.Equal(t, 3000, getHighestNumber([]int{1200, 3000, 2300, 1700}))
}

func getHighestNumber(numbers []int) int {
	sort.Ints(numbers)

	if len(numbers) > 0 {
		return numbers[len(numbers)-1]
	}

	return 0
}

func TestUpdateStringInSlices(t *testing.T) {
	slices := []string{"dia", "suka", "membaca"}
	old := "membaca"
	new := "makan"

	updateStringInSlices(slices, old, new)

	assert.Equal(t, []string{"dia", "suka", "makan"}, slices)
}

func updateStringInSlices(slices []string, old, new string) {
	for k, v := range slices {
		if v == old {
			slices[k] = new
			break
		}
	}
}

func TestGetAbsoluteValue(t *testing.T) {
	assert.Equal(t, 10, getAbsoluteValue(10))
	assert.Equal(t, 20, getAbsoluteValue(-20))
}

func getAbsoluteValue(number int) int {
	if number < 0 {
		return -number
	}

	return number
}

func TestGetPercentage(t *testing.T) {
	assert.Equal(t, "25.00%", getPercentage(50, 200))
}

func getPercentage(n1, n2 float64) string {
	result := fmt.Sprintf("%.2f", n1*100/n2) + "%"
	return result
}
