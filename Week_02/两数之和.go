
// 两数相加暴力求解
func twoSum(nums []int, target int) []int {
    len := len(nums)
    var result []int
    for i := 0; i < len - 1; i++ {
        for j := i + 1; j < len; j++ {
            if (nums[i] + nums[j]) == target {
                result = append(result, i)
                result = append(result, j)
            }
        }
    }
    return result
}

// 一遍哈希法 时间复杂度O(n) 空间复杂度O(n)
func twoSum(nums []int, target int) []int {
    hash := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        if _,ok := hash[nums[i]]; ok {
            return []int{hash[nums[i]], i}
        } else {
            diff := target - nums[i]
            hash[diff] = i
        }
    }
    return []int{}
}