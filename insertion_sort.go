package algorithms

var insert = func(array []int, rightIndex, value int) []int {
	j := rightIndex
	for j >= 0 && array[j] > value {
		array[j + 1] = array[j]
		j--
	}
	array[j + 1] = value
	return array
}

var InsertionSort = func(array []int, startPercentage int, endPercentage int) []int {
	startPoint := len(array) * startPercentage / 100
	endPoint := len(array) * endPercentage / 100

	for currentIndex := startPoint + 1; currentIndex < endPoint; currentIndex++ {
		insert(array, currentIndex - 1, array[currentIndex]);
	}
	return array
}
