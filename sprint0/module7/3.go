package main 

func SliceCopy(nums []int) []int{
    copy := make([]int, len(nums))
    copy = append([]int{}, nums...)
    return copy
}