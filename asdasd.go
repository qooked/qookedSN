package main

func moveZeroes(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
		}
	}
	return nums
}
