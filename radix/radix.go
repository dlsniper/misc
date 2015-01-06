package radix

import "math"

func Radix(arr []int) []int {
	noElements := len(arr)

	if noElements < 2 {
		return arr
	}

	var result []int

	if noElements == 2 {
		if arr[0] > arr[1] {
			result = append(result, arr[1], arr[0])
		} else {
			result = arr
		}

		return result
	}

	return arr
}

func ClosestTwoPower(x int) int {
	return int(math.Ceil(math.Log2(float64(x))))
}
