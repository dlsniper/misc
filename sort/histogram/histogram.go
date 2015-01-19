package histogram

import "math"

var temp = []bool{}

func Init() {
	if len(temp) == 0 {
		for i := 0; i < math.MaxInt16*2; i++ {
			temp = append(temp, false)
		}
	} else {
		for i := 0; i < math.MaxInt16*2; i++ {
			temp[i] = false
		}
	}
}

func HistogramSort(array []int) []int {
	Init()
	return histogramSort(array)
}

func histogramSort(array []int) []int {
	result := []int{}

	for _, elem := range array {
		temp[elem] = true
	}

	for key := range temp {
		if temp[key] {
			result = append(result, key)
		}
	}

	return result
}
