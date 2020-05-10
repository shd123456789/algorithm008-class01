// 2遍
// 分治法时间复杂度为O(nlogn) 空间为O(logn)
func majorityElement(nums []int) int {
    return helper(nums, 0, len(nums) - 1)
}

func helper(nums[] int, lp int, rp int) int {
    if lp == rp {
        return nums[lp]
    }
    mid := (rp - lp)/2 + lp
    left := helper(nums, lp, mid)
    right := helper(nums, mid + 1, rp)
    if left == right {
        return left
    }
    leftCount := count(nums, lp, rp, left)
    rightCout := count(nums, lp, rp, right)
    if leftCount > rightCout {
        return left
    } else {
        return right
    }
}

func count(nums[] int, lp int, rp int, val int) int {
    count := 0
    for i := lp; i <= rp; i++ {
        if nums[i] == val {
            count++
        }
    }
    return count
}

// 投票法时间为O(n),空间为O(1)
func majorityElement(nums []int) int {
    count,major := 0,0
    for _,v := range nums {
        if count == 0 {
            major = v
            count++
        } else if v == major {
            count++
        } else {
            count--
        }
    }
    return major
}
