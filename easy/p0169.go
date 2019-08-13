package easy

func majorityElement(nums []int) int {
	l := len(nums)
	lHalf := l / 2
	counter := make(map[int]int)
	for i := 0; i < l; i++ {
		num := nums[i]
		counter[num]++
		if counter[num] > lHalf {
			return num
		}
	}
	return -1
}
