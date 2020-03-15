package longestPalindromicSubstring

import "strings"

func longestPalindrome(s string) string {
	str:=strings.Split(s, "")
	if len(str)<=1{
		return s
	}

	dp:=make([][]int,len(str))
	for k,v:=range dp{
		v=make([]int, len(str))
		dp[k]=v
	}
	for i:=0;i<len(str);i++{
		for j:=0;j<=i;j++{
			find(j,i,str,&dp)
		}
	}
	//fmt.Println(dp)
	max:=0
	maxstr:=[]string{}
	for i:=0;i<len(str);i++{
		for j:=0;j<len(str);j++{
			if dp[i][j]==1{
				if j-i+1 > max{
					maxstr=str[i:j+1]
					max=j-i+1
				}
			}
		}
	}

	return strings.Join(maxstr, "")
}

func find(i,j int, str []string, dp *[][]int){
	if i==j{
		(*dp)[i][j]=1
	}else if str[i]==str[j]{
		if i+1==j{
			(*dp)[i][j]=1
		}else{
			if (*dp)[i+1][j-1]==1{
				(*dp)[i][j]=1
			}
		}
	}

}

func reserver(s []string) []string{
	s2:=[]string{}
	for i:=len(s)-1;i>=0; i-- {
		s2=append(s2, s[i])
	}
	return s2
}