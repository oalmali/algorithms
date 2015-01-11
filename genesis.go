package algorithms

import (
	"math/rand"
	"time"
	"math"
)

//reference:wikipedia.org/wiki/Power_of_10
var powersOfTen = map[string]int  {
	"one" : 0,
	"ten" : 1,
	"hundred" : 2,
	"thousand" : 3,
	"ten thousand" : 4,
	"hundred thousand" : 5,
	"million" : 6,
	"billion" : 9,
	"trillion" : 12,
	"quadrillion" : 15,
	"sextillion" : 21,
	"septillion" : 24,
	"octillion" : 27,
	"nonillion" : 30,
	"decillion" : 33,
	"undecillion" : 36,
	"duodecillion" : 39,
	"tredecillion" : 42,
	"quattordecillion" : 45,
	"quindecillion" : 48,
}

func Genesis(size string) []int {
	pow, ok := powersOfTen[size]
	if !ok {
		return nil
	}
	limit := int(math.Pow(10, float64(pow)))
	array := []int{};
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < limit; i++ {
		array = append(array, rand.Intn(limit))
	}
	return array
}
