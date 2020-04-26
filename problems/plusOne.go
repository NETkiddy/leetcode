package problems

func plusOne(digits []int) []int {
	if len(digits)==0{
		return []int{1}
	}
	step:=1
	for i:=len(digits)-1;i>=0;i--{
		if digits[i]+step>=10{
			digits[i]=digits[i]+step-10
			step=1
		}else{
			digits[i]=digits[i]+step
			step=0
		}
	}
	if step==1{
		digits = append([]int{step}, digits...)
	}
	return digits
}
