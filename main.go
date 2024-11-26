package main

import (
	"fmt"
	"math"
	"sort"
	// "guess-it-1/student/functions"
)

func Average(values []int) float64 {
	var total float64

	for _, v := range values {
		total += float64(v)
	}

	return total / float64(len(values))
}

func Median(values []int) float64 {
	var result float64
	sort.Ints(values)
	if len(values)%2 == 0 {
		result = math.Round((float64(values[len(values)/2]) + float64(values[(len(values)/2)-1])) / 2)
	} else {
		result = float64(values[len(values)/2])
	}
	return result
}

func Variance(values []int) float64 {
	var result float64

	for _, v := range values {
		result += (float64(v) - Average(values)) * (float64(v) - Average(values))
	}
	result = result / float64(len(values))
	return math.Round(result)
}

func StandardDeviation(Variance float64) float64 {
	return math.Sqrt(Variance) + 0.5
}

func CalculateRange(checkingValues []int) (int, int) {
	if len(checkingValues) < 2 {
		return (checkingValues[len(checkingValues)-1] - 50), (checkingValues[len(checkingValues)-1] + 50)
	}

	lowerLimit := int(Average(checkingValues) - 1.5*StandardDeviation(Variance(checkingValues)))
	upperLimit := int(Average(checkingValues) + 1.5*StandardDeviation(Variance(checkingValues)))

	return lowerLimit, upperLimit
}

func main() {
	var values []int

	for i := 0; i <= 12500; i++ {
		var n int
		fmt.Scanf("%d", &n)
		values = append(values, n)
		v1, v2 := CalculateRange(values)
		fmt.Printf("%d %d", v1, v2)
	}
}
