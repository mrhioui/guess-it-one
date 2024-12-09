package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

// Function to calculate the average of a slice of integers
func Average(values []int) float64 {
	var total float64
	for _, v := range values {
		total += float64(v)
	}
	return total / float64(len(values))
}

// Function to calculate the median of a slice of integers
func Median(values []int) float64 {
	sort.Ints(values)
	n := len(values)
	if n%2 == 0 {
		return (float64(values[n/2-1]) + float64(values[n/2])) / 2
	}
	return float64(values[n/2])
}

// Function to calculate the variance of a slice of integers
func Variance(values []int) float64 {
	avg := Average(values)
	var total float64
	for _, v := range values {
		total += (float64(v) - avg) * (float64(v) - avg)
	}
	return total / float64(len(values))
}

// Function to calculate the standard deviation given the variance
func StandardDeviation(variance float64) float64 {
	return math.Sqrt(variance)
}

// Function to calculate the lower and upper limits for outliers
func CalculateRange(values []int) (int, int) {
	if len(values) < 2 {
		return values[len(values)-1] - 50, values[len(values)-1] + 50
	}
	avg := Average(values)
	stdDev := StandardDeviation(Variance(values))
	lowerLimit := int(avg - 1.5*stdDev)
	upperLimit := int(avg + 1.5*stdDev)
	return lowerLimit, upperLimit
}

func main() {
	var n int
	var values []int

	fmt.Println("Enter values, followed by 'Ctrl+D' to stop:")

	// Loop to accept values from the user
	for {
		_, err := fmt.Fscan(os.Stdin, &n)
		if err != nil {
			if err.Error() == "EOF" { // End of input
				break
			}
			os.Exit(1)
		}
		values = append(values, n)

		if len(values) > 1 {
			// Calculate and print the range limits
			v1, v2 := CalculateRange(values)
			fmt.Printf("%d %d\n", v1, v2)
		}
	}
}
