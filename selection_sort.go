package algorithms

var swap = func(array []int, firstIndex, secondIndex int) {
	var temp = array[firstIndex];
	array[firstIndex] = array[secondIndex];
	array[secondIndex] = temp;
}

var indexOfMinimum = func(array []int, startIndex int, endIndex int) int {
	minValue := array[startIndex];
	minIndex := startIndex;

	for i := minIndex + 1; i < endIndex; i++ {
		if(array[i] < minValue) {
			minIndex = i;
			minValue = array[i];
		}
	}
	return minIndex;
}

var SelectionSort = func(array []int, startPercentage int, endPercentage int) []int {
	startPoint := len(array) * startPercentage / 100
	endPoint := len(array) * endPercentage / 100

	var indexOfCurrent = 0;
	for	i := startPoint; i < endPoint; i++ {
		indexOfCurrent = indexOfMinimum(array, i, endPoint);
		swap(array, indexOfCurrent, i);
	}
	return array
}
