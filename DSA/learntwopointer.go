package main

import "fmt"

func isPalindrome(arr []int) bool {
	left, right := 0, len(arr)-1

	for left < right {
		if arr[left] != arr[right] {
			return false
		}
		left++
		right--
	}
	return true

}
func hasPairWithSum(arr []int, target int) bool {
	left, right := 0, len(arr)-1
	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			return true
		}
		if sum < target {
			left++
		} else {
			right--
		}

	}
	return false
}
func removeDuplicates(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	l := 0
	for r := 1; r < len(arr); r++ {
		if arr[l] != arr[r] { //
			l++
			arr[l] = arr[r] //
		}
	}
	return l + 1
}

func main() {
	fmt.Println(isPalindrome([]int{1, 2, 3, 2, 1}))
	fmt.Println(isPalindrome([]int{1, 2, 3, 4, 5}))
	fmt.Println(hasPairWithSum([]int{1, 2, 3, 4, 5}, 5))
	fmt.Println(removeDuplicates([]int{1, 2, 3, 2, 1}))
}
