package easy

func TitleToNumber(s string) int {
	var sum int
	for _, k := range s {
		tmp := k - '@'
		sum = sum*26 + int(tmp)
	}
	return sum
}
