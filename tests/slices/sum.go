package main

func Sum(arr []int) int {
	length := len(arr)
	sum := 0

	for i := 0; i < length; i++ {
		sum = sum + arr[i]
	}
	
	return sum
}