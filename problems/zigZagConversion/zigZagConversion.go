package zigZagConversion

import (
	"strings"
)

// by filling matrix
func convert(str string, numRows int) string {
	if str=="" || numRows<=1{
		return str
	}
	s:=strings.Split(str,"")
	data:="";
	batch:=len(s)/(numRows+numRows-2)
	left := len(s)%(numRows+numRows-2)-numRows+1
	if len(s)%(numRows+numRows-2)-numRows<0{
		left = 1
	}
	col:=(numRows-1)*batch+left
	res:=make([][]string,col*numRows)
	for k:=0;k<=numRows-1;k++{
		res[k]=make([]string,col)
	}
	i,j,count:=0,0,0
	isUp:=true
	for {
		if count == len(s){
			break
		}
		res[i][j]=s[count]
		count++

		if isUp{
			i++
		}else{
			i--
			j++
		}
		if i== numRows-1{
			isUp=false
		}
		if i == 0{
			isUp=true
		}
	}

	for _,row:=range res{
		for _,v:=range row{
			if v!=""{
				data+=v
			}
		}
	}
	return data
}
