func getKthMagicNumber(k int) int {
    if k <= 0 {
        return 0
    }   
    nums := make([]int, k) 
    nums[0] = 1
    p1,p2,p3 := 0,0,0
    for i := 1; i < k; i++ {
        nums[i] = min(nums[p1] * 3, nums[p2] * 5, nums[p3] * 7)
        if nums[i] == nums[p1] * 3 {
            p1++
        }
        if nums[i] == nums[p2] * 5 {
            p2++
        }
        if nums[i] == nums[p3] * 7 {
            p3++
        }
    }
    return nums[k - 1]
}

func min(n1 int,n2 int,n3 int) int {
    var min int
    if n1 > n2 {
        min = n2
    } else {
        min = n1
    }
    if min > n3 {
        min = n3
    }
    return min
}
