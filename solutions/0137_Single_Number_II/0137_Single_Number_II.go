package leetcode

func singleNumber(nums []int) int {
	var ans int32
	for i := 0; i < 32; i++ {
		count := 0
		for _, x := range nums {
			if (x>>i)&1 == 1 {
				count += 1
			}
		}
		if count%3 != 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}
