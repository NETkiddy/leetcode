package problems

import "fmt"

func isOneBitCharacter(bits []int) bool {
	if len(bits)==0{
		return false
	}
	res := false
	for i:=0;i<len(bits);{
		if bits[i]==0{
			i++
			res=true
		}else if bits[i]==1{
			i+=2
			res=false
		}else{
			fmt.Println("error")
		}
	}
	return res
}