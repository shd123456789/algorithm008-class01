
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

// 一遍哈希法
func twoSum(nums []int, target int) []int {
    len := len(nums)
    var hash = make(map[int]int)
    for i := 0; i <= len - 1; i++ {
        if i > 0 { // 判断是否存在对应相加的数的哈希
            diff := target - nums[i]
            if _,ok := hash[diff]; ok {
                return []int{hash[diff], i}
            }
        }
        hash[nums[i]] = i; 
    }
    return []int{}
}