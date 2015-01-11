package algorithms

import "math"

var doSearch = func(array []int, targetValue int) int {
	var min = 0;
	var max = len(array) - 1;
	var guess = 0;
	for min <= max {
		guess = int(math.Floor(float64((max + min) / 2)));
		var currentElement = array[guess];
		if currentElement == targetValue {
			return guess;
		} else if currentElement < targetValue {
			min++;
		} else {
			max--;
		}
	}
	return -1
}
