package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	arr := [][]int{}

	if len(nums) == 0 || (len(nums) == 1 && nums[0] == 0) {
		return [][]int{}
	}

	for i := 0; i < len(nums)-2; i++ {
		start := i + 1
		end := len(nums) - 1

		for start < end {
			sum := nums[i] + nums[start] + nums[end]

			switch {
			case sum == 0:
				temp := []int{nums[i], nums[start], nums[end]}
				arr = append(arr, temp)
				for start < len(nums)-1 && nums[start] == nums[start+1] {
					start += 1
				}

				for end < len(nums)-1 && nums[end] == nums[end-1] {
					end -= 1
				}
				start += 1
				end -= 1
			case sum > 0:
				end -= 1
			case sum < 0:
				start += 1
			}

			for i < len(nums)-2 && nums[i] == nums[i+1] {
				i += 1
			}
		}
	}
	return arr
}

func main() {
	fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
	fmt.Println(threeSum([]int{0,0,0}))
}

//Time complexity: O(n^2)
//Runtime: 28 ms
//Memory usage: 7 MB
