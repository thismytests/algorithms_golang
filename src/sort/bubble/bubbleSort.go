package bubble

import "fmt"

func BubbleSort(arr []int) {
	for i := len(arr); i > 0; i-- {
		for j := 1; j < i; j++ {

			if arr[j-1] > arr[j] {
				buffer := arr[j]

				arr[j] = arr[j-1]
				arr[j-1] = buffer
			}
		}
	}

	fmt.Println("", arr)
}
