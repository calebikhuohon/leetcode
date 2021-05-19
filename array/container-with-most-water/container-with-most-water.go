package main

import "fmt"

func maxArea(height []int) int {
	p1 := 0
	p2 := len(height) - 1
	max := 0

	for p1 < p2 {
		var area int

		if height[p1] < height[p2] {
			area = height[p1] * (p2 -p1)
			p1 += 1
		} else {
			area = height[p2] * (p2 -p1)
			p2 -= 1
		}

		if area > max {
			max = area
		}

	}
	return max
}

func main()  {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}
