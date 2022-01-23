package _select

func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		var indexMin int = i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[indexMin] {
				indexMin = j
			}
		}

		tmp := arr[i]
		arr[i] = arr[indexMin]
		arr[indexMin] = tmp
	}
}
