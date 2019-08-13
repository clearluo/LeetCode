package easy

func Generate(numRows int) [][]int {
	retData := [][]int{}
	tmp := 0
	for i := 0; i < numRows; i++ {
		if i == 0 {
			retData = append(retData, []int{1})
			continue
		}
		item := []int{}
		tmp = 0
		for _, v := range retData[i-1] {
			item = append(item, tmp+v)
			tmp = v
		}
		item = append(item, 1)
		retData = append(retData, item)
	}
	return retData
}
