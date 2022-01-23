package binar

func BinarySearch(array []int, target int) int {
	startIndex := 0
	endIndex := len(array) - 1
	midIndex := len(array) / 2

	for startIndex <= endIndex {

		middleValue := array[midIndex]

		if middleValue == target {
			return midIndex
		}

		if middleValue > target {
			endIndex = midIndex - 1
		} else {
			startIndex = midIndex + 1
		}

		midIndex = (startIndex + endIndex) / 2

	}

	return -1
}
