package sort

//三路快排
func ThreeWayFastDischarge(nums []int) {
	zero := -1
	two := len(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 2 {
			two = i
		} else {
			break
		}
	}
	index := zero + 1
	for index < two {
		num := nums[index]
		if num == 0 {
			zero = zero + 1
			swap(nums, zero, index)
			index++
		} else if num == 2 {
			two = two - 1
			swap(nums, index, two)
			if nums[index] == 1 {
				index++
			} else if index != two && nums[index] == 0 {
				zero = zero + 1
				swap(nums, index, zero)
				index++
			} else if index != two && nums[index] == 2 {
				two = two - 1
				swap(nums, index, two)
			}
		} else if num == 1 {
			index++
		}
	}
}

func swap(nums []int, index1 int, index2 int) {
	temp := nums[index1]
	nums[index1] = nums[index2]
	nums[index2] = temp
}
