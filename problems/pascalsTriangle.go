package problems

func generate(numRows int) [][]int {
	if numRows==0{
		return nil
	}
	res := make([][]int,0)
	for i:=0;i<numRows;i++{
		res=append(res, make([]int,i+1))
	}
	for i:=0;i<numRows;i++{
		for j:=0;j<=i;j++{
			if j==0||j==len(res[i])-1{
				res[i][j]=1
			}else{
				res[i][j]=res[i-1][j-1]+res[i-1][j]
			}
		}
	}
	return res
}