package easy

func hammingWeight(num uint32) int {
	count := 0
	for num > 1 {
		if num%2 == 1 {
			count++
		}
		num /= 2
	}
	if num == 1 {
		count++
	}
	return count
}
