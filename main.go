package main

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	temp := []int{0, 0, 0}
	combination(len(arr), len(temp), 0, 0, arr, temp)
}
