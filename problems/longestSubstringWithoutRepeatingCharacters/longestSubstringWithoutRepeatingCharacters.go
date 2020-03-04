package longestSubstringWithoutRepeatingCharacters

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(str string) int {
	s:=strings.Split(str, "")
	fmt.Println(s)
	max:=make([]int,len(s))
	for i:=len(s)-1;i>=0;i--{
		if i==len(s)-1{
			max[i]=1
			continue
		}
		if s[i]==s[i+1]{
			max[i]=1
			continue
		}
		if s[i]!=s[i+1]{
			start:=i+1
			end:=i+1+max[i+1]
			pos:=contains(s[start:end], s[i])
			if pos==-1{
				max[i]=max[i+1]+1
			}else{
				max[i]=pos+1
			}
			continue
		}
	}
	res:=0
	for _,v:=range max{
		fmt.Println(v)
		if v>res{
			res=v
		}
	}
	return res
}


func contains(s []string, sub string) int{
	index:=-1
	for k,v:=range s{
		if v==sub{
			index=k
			break
		}
	}
	return index
}