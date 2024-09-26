package main 

func Clean(nums []int, x int) []int {
    if len(nums) == 0 {
        return nums
    }

    i := 0
    for j := 0; j < len(nums); j++ {
        if nums[j] != x {
            nums[i] = nums[j]
            i++
        }
    }
    nums = nums[:i]
    return nums
}