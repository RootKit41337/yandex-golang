package main

func Mix(nums []int) []int {
    n := len(nums) / 2
    result := make([]int, len(nums))
    for i := 0; i < n; i++ {
        result[2*i] = nums[i]
        result[2*i+1] = nums[n+i]
    }
    return result
}