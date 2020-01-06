package utils

// []int {9,8,7,6,5}
// []int {5,6,7,8,9}

//BubbleSort ...
func BubbleSort(elements []int) []int {
	isKeepRunning := true
	for isKeepRunning {
		isKeepRunning = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				isKeepRunning = true
			}
		}
	}

	return elements
}
