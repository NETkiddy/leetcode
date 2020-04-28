package problems

//需要利用最大公约数
func hasGroupsSizeX(deck []int) bool {
	set:=make(map[int]int)
	for _,d:=range deck{
		set[d]++
	}
	//fmt.Println(set)
	res:=0
	for _,c:=range set{
		//fmt.Println(k,c,res)
		res = gcd(c,res)
		//fmt.Println(res)
	}

	if res >1{
		return true
	}
	return false
}

func gcd(a,b int) int{
	if b==0{
		return a
	}
	if a==b{
		return a
	}
	if a<b{
		t:=a
		a=b
		b=t
	}
	return gcd(b,a%b)
}