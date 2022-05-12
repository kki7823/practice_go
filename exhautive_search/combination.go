package exhautive_search

import "fmt"

func combination(N int, R int, n int, r int, arr []int, temp []int) {
	if r == R {
		fmt.Println(temp)
		return
	}

	if n == N {
		return
	}

	temp[r] = arr[n]

	combination(N, R, n+1, r+1, arr, temp)
	combination(N, R, n+1, r, arr, temp)
}