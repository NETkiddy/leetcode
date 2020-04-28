package problems
//通过辅助空间来加速寻找
func fairCandySwap(A []int, B []int) []int {
	res:=make([]int,0)
	ta:=sum(A)
	tb:=sum(B)
	if ta==tb{
		return res
	}


	seta:=make(map[int]int,0)
	for _,a:=range A{
		seta[a]=1
	}
	if ta>tb{
		dif:=ta-tb
		for _,b:=range B{
			if _,found:=seta[b+dif/2];found{
				res=append(res,b+dif/2)
				res=append(res,b)
				return res
			}
		}
	}
	if ta<tb{
		dif:=tb-ta
		for _,b:=range B{
			if _,found:=seta[b-dif/2];found{
				res=append(res,b-dif/2)
				res=append(res,b)
				return res
			}
		}
	}

	return res
}

func sum(arr []int) int{
	t:=0
	for _,a:= range arr{
		t+=a
	}
	return t
}