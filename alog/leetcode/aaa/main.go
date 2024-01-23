package main

import "fmt"

//给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，
//找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
var ans [][]int
func baiduTest(nums []int,target int){
	for i := 0;i<len(nums);i++{
		view := make([]bool,len(nums))
		view[i] = true
		recorder := []int{nums[i]}
		recuricv(nums,view,recorder,target-nums[i])
	}
}

func recuricv(nums []int,view []bool,recorder []int,remaining int){
	if remaining <0{
		return
	}
	if remaining == 0{
		ans = append(ans,recorder)
		return
	}
	for i := 0;i<len(nums);i++{
		if !view[i]{

			newView := make([]bool,len(nums))
			copy(newView,view)
			newView[i] = true

			var newRecorder []int
			for j := 0;j<len(recorder);j++{
				newRecorder = append(newRecorder,recorder[j])
			}

			recuricv(nums,newView,newRecorder,remaining-nums[i])
		}
	}
}


func main() {
	nums := []int{1,2,3,4,6,7}
	target := 7
	baiduTest(nums,target)
	for _,item := range ans{
		fmt.Println(item)
	}
}
