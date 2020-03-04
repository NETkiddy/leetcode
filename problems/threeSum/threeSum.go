package threeSum

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3{
		return [][]int{}
	}
	sort.Ints(nums)
	res := make([][]int,0)
	for k:=0;k<len(nums);{
		c := -1*nums[k]
		i:=k+1
		j:=len(nums)-1
		for ;i<j;{
			if nums[i]+nums[j]==c{
				res = append(res, []int{nums[k],nums[i],nums[j]})
				fmt.Printf("i:%d,j:%d",i,j)
				for ;i<j;{
					i++
					if nums[i]!=nums[i-1]{
						break
					}
				}
				for ;i<j;{
					j--
					if nums[j]!=nums[j+1]{
						break
					}
				}
			}else if nums[i]+nums[j]<c{
				i++
			}else{
				j--
			}
		}
		for ;;{
			k++
			if k==len(nums) || nums[k] != nums[k-1]{
				break
			}
		}
	}
	return res
}