package containerWithMostWater

/*
//brute
func maxArea(height []int) int {
    max := 0
    for i:=0;i<len(height)-1;i++{
        for j:=i+1;j<len(height);j++{
            high := height[i]
            if height[j] < high{
                high = height[j]
            }
            if  (j-i)*high > max {
                max = (j-i)*high
            }
        }
    }

    return max
}
*/


// two pointers
func maxArea(height []int) int {
    i:=0
    j:=len(height)-1
    max := 0

    for ;i<j; {
        long:=j-i
        high := height[i]
        if height[j] < height[i]{
            high = height[j]
            j--
        }else{
            high = height[i]
            i++
        }
        if  long*high > max {
            max = long*high
        }
    }

    return max
}

