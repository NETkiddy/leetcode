package integerToRoman

//just to demo the algorithm
func intToRoman(num int) string {
	res:=""
	cnt:=num/1000
	left:=num%1000
	for i:=0;i<=cnt-1;i++{
		res+="M"
	}
	cnt=left/(1000-100)
	if cnt>=1{
		for i:=0;i<=cnt-1;i++{
			res+="CM"
		}
		left=left%(1000-100)
	}
	cnt=left/500
	if cnt>=1{
		for i:=0;i<=cnt-1;i++{
			res+="D"
		}
		left=left%500
	}
	cnt=left/(500-100)
	if cnt>=1{
		for i:=0;i<=cnt-1;i++{
			res+="CD"
		}
		left=left%(500-100)
	}
	cnt=left/100
	if cnt>=1{
		for i:=0;i<=cnt-1;i++{
			res+="C"
		}
		left=left%100
	}
	cnt=left/(100-10)
	if cnt>=1{
		for i:=0;i<=cnt-1;i++{
			res+="XC"
		}
		left=left%(100-10)
	}
	cnt=left/50
	if cnt>=1{
		for i:=0;i<=cnt-1;i++{
			res+="L"
		}
		left=left%50
	}
	cnt=left/(50-10)
	if cnt>=1{
		for i:=0;i<=cnt-1;i++{
			res+="XL"
		}
		left=left%(50-10)
	}
	cnt=left/10
	if cnt>=1{
		for i:=0;i<=cnt-1;i++{
			res+="X"
		}
		left=left%10
	}
	cnt=left/(10-1)
	if cnt>=1{
		for i:=0;i<=cnt-1;i++{
			res+="IX"
		}
		left=left%(10-1)
	}
	cnt=left/5
	if cnt>=1{
		for i:=0;i<=cnt-1;i++{
			res+="V"
		}
		left=left%5
	}
	cnt=left/(5-1)
	if cnt>=1{
		for i:=0;i<=cnt-1;i++{
			res+="IV"
		}
		left=left%(5-1)
	}
	cnt=left/1
	if cnt>=1{
		for i:=0;i<=cnt-1;i++{
			res+="I"
		}
	}
	return res
}
