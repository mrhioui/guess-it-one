package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

func Median(values []int) float64 {
	sort.Ints(values)
	n := len(values)
	if n%2 == 0 {
		return (float64(values[n/2-1]) + float64(values[n/2])) / 2
	}
	return float64(values[n/2])
}

func CalculateRange(values []int) (int, int) {
	median := Median(values)
	cons := 40.0
	lowerLimit := median - cons
	upperLimit := median + cons
	return int(math.Round(lowerLimit)), int(math.Round(upperLimit))
}

func main() {
	var n int
	var values []int
	for {
		_, err := fmt.Fscan(os.Stdin, &n)
		if err != nil {
			fmt.Println(err)
			if err.Error() == "EOF" {
				break
			}
			os.Exit(1)
		}
		values = append(values, n)

		if len(values) > 1 {
			v1, v2 := CalculateRange(values)
			fmt.Println(v1, v2)
		}
	}
}
