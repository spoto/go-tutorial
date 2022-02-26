package main

import "fmt"

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("cannot compute the maximum of zero elements")
	}
	max := vals[0]
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max, nil
}

func max2(first int, vals ...int) int {
	max := first
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}
