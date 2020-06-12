package leetcode

func longestMountain(A []int) int {
	N := len(A)
	ans := 0
	left := 0
	for left < N {
		right := left
		// check if left boundary
		if right+1 < N && A[right] < A[right+1] {
			for right+1 < N && A[right] < A[right+1] {
				right++
			}

			// check if right is the peak
			if right+1 < N && A[right] > A[right+1] {
				// move right to right boundary
				for right+1 < N && A[right] > A[right+1] {
					right++
				}
				ans = Max(ans, right-left+1)
			}
		}
		left = Max(left+1, right)
	}
	return ans
}

func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
