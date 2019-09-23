package easy

func MissingNumber(nums []int) int {
	l := len(nums)
	sum := l * (l + 1) / 2
	for _, v := range nums {
		sum -= v
	}
	return sum
}
