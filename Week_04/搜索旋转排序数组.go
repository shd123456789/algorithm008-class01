// 时间复杂度为O(logn) 空间复杂度为O(1)
func search(nums []int, target int) int {
    l,r := 0,len(nums) - 1
    for l <= r {
        mid := l + (r - l) >> 2
        if nums[mid] == target {
            return mid
        }
        if nums[mid] >= nums[l] {
            if target >= nums[l] && target < nums[mid] {
                r = mid - 1
            } else {
                l = mid + 1
            }
        } else {
            if target > nums[mid] && target <= nums[r] {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
    }
    return -1
}