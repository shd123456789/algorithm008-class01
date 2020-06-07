//15. 三数之和(3遍) 时间复杂度O(n^2)空间复杂度为O(1)(排除结果的空间)
func threeSum(nums []int) [][]int {
    var results [][]int
    sort.Ints(nums)
    for i := 0; i < len(nums) - 2; i++ {
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }
        target,left,right := -nums[i], i + 1, len(nums) - 1
        for left < right {
            sum := nums[left] + nums[right]
            if target == sum {
                results = append(results, []int{nums[i], nums[left], nums[right]})
                left++
                right--
                for left < right && nums[left] == nums[left - 1] {
                    left++
                }
                for left < right && nums[right] == nums[right + 1] {
                    right--
                }
            } else if target > sum {
                left++
            } else if target < sum {
                right--
            }
        }
    }
    return results
}