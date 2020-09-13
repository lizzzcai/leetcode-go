package leetcode

import "fmt"

var max_time int

func largestTimeFromDigits(A []int) string {
	max_time = -1
	permutate(A, 0)
	if max_time == -1 {
		return ""
	} else {
		return fmt.Sprintf("%02d:%02d", max_time/60, max_time%60)
	}
}

func build_time(permutation []int) {
	h, i, j, k := permutation[0], permutation[1], permutation[2], permutation[3]
	hour := h*10 + i
	minute := j*10 + k
	if hour < 24 && minute < 60 {
		time := hour*60 + minute
		if time > max_time {
			max_time = time
		}
	}
}

func permutate(arr []int, start int) {
	if start == len(arr) {
		build_time(arr)
		return
	}

	for idx := start; idx < len(arr); idx++ {
		// swap the start and index
		arr[start], arr[idx] = arr[idx], arr[start]
		permutate(arr, start+1)
		// swap back
		arr[start], arr[idx] = arr[idx], arr[start]
	}
}
