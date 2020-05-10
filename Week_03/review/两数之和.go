// 2遍
func twoSum(nums []int, target int) []int {
    len := len(nums)
    var hash = make(map[int]int)
    for i := 0; i <= len - 1; i++ {
        diff := target - nums[i]
        if _,ok := hash[diff]; ok {
            return []int{hash[diff], i}
        }
        hash[nums[i]] = i; 
    }
    return []int{}
}