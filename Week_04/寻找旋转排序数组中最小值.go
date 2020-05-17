// 时间复杂度为O（logn） 空间为O(1)
func findMin(nums []int) int {
    l,r := 0,len(nums) - 1
    minV := math.MaxInt64
    for l <= r {
        mid := l + (r - l) >> 2
        if nums[mid] >= nums[l] {
            minV = min(minV, nums[l])
            l = mid + 1
        } else {
            minV = min(minV, nums[mid])
            r = mid - 1
        }
    }
    return minV
}

func min(n1 int, n2 int) int {
    if n1 > n2 {
        return n2
    }
    return n1
}