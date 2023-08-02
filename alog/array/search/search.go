package search

//二分查找
func BinarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	left := 0
	right := len(nums) - 1
	lastmiddle := 0
	for {
		if left == right-1 {
			if nums[left] == target {
				return left
			}
			if nums[right] == target {
				return right
			}
			return -1
		}
		middle := (right + left) / 2
		if lastmiddle == middle {
			return -1
		}
		if target < nums[middle] {
			right = middle
		}
		if target > nums[middle] {
			left = middle
		}
		if target == nums[middle] {
			return middle
		}
		lastmiddle = middle
	}
}
