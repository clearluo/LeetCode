package easy

func MoveZeroes(nums []int) {
	l := len(nums)
	for i := l - 1; i >= 0; i-- {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
		}
	}
	//for i, v := range nums {
	//	if i == 0 {
	//		fmt.Printf("%s%d", "[", v)
	//	} else {
	//		fmt.Printf(",%d", v)
	//	}
	//}
	//fmt.Println("]")
}
