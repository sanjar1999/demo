package main

import "fmt"

//fibonachchi

func fib(num int) int {
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	} else {
		return fib(num-1) + fib(num-2)
	}
}

//palindrom

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

//odd even sum

func OddEvenSum(input int) (int, int) {
	odd := 0
	even := 0
	for i := 0; i <= input; i++ {
		if i%2 == 0 {
			odd += i
		}
		if i%2 != 0 {
			even += i
		}
	}
	return odd, even
}

// has duplicate array

func hasDuplicate(arr []int) int {
	visited := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		if visited[arr[i]] {
			return arr[i]
		} else {
			visited[arr[i]] = true
		}
	}
	return -1
}

func main() {

	//fizzbuzz
	fizz := 10
	buzz := 20
	given := 40

	if given%fizz == 0 && given%buzz == 0 {
		fmt.Println("fizz buzz")
	} else if given%buzz == 0 {
		fmt.Println("buzz")
	} else if given%fizz == 0 {
		fmt.Println("fizz")
	}

	fmt.Println(fib(5))
	fmt.Println(isPalindrome("anna"))
	fmt.Println(OddEvenSum(32))
	fmt.Println(hasDuplicate([]int{1, 5, 2, 3, 4, 5, 6}))
}
