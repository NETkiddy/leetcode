package twoSum

// one pass map
func twoSum(nums []int, target int) []int {
	rmap:=make(map[int]int)

	for i,v := range nums{
		left := target - v
		if k,found:= rmap[left];found{
			if k !=i{
				return []int{k, i}
			}
		}
		rmap[v] = i
	}

	return nil
}

/*
// two pass map
func twoSum(nums []int, target int) []int {
    rmap:=make(map[int]int)
    for i, v:= range nums{
       rmap[v]=i
    }

    for i,v := range nums{
        left := target-v
        if k,found:=rmap[left];found{
            if k!=i{
                return []int{i,k}
            }
        }
    }
    return nil

}
*/