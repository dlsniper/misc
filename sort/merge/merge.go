package merge

func MergeSort(array []int) []int {
	mergeSort(&array, 0, len(array)-1)

	return array
}

func mergeSort(array *[]int, low, high int) {
	if low < high {
		middle := (low + high) / 2

		mergeSort(array, low, middle)
		mergeSort(array, middle+1, high)
		merge(array, low, middle, high)
	}
}

func merge(array *[]int, low, middle, high int) {
	helper := map[int]int{}
	for i := low; i <= high; i++ {
		helper[i] = (*array)[i]
	}

	helperLeft := low
	helperRight := middle + 1
	current := low

	for helperLeft <= middle && helperRight <= high {
		if helper[helperLeft] <= helper[helperRight] {
			(*array)[current] = helper[helperLeft]
			helperLeft++
		} else {
			(*array)[current] = helper[helperRight]
			helperRight++
		}
		current++
	}

	remaining := middle - helperLeft
	for i := 0; i <= remaining; i++ {
		(*array)[current+i] = helper[helperLeft+i]
	}
}
