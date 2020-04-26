func numberOfSubarrays(nums []int, k int) int {
    var oddNumbers []int  
    num := 0
    oddNumbers = append(oddNumbers, -1)
    for i := 0; i < len(nums); i++ {
        if (nums[i] % 2) != 0 {
            oddNumbers = append(oddNumbers, i) // 记录下奇数位置
        } 
    } 
    oddNumbers = append(oddNumbers, len(nums))
    if k > len(oddNumbers) {
        return 0
    }
    for i := 1; i < len(oddNumbers) - k; i++ {
        num += (oddNumbers[i] - oddNumbers[i - 1]) * (oddNumbers[k + i] - oddNumbers[k + i - 1])
    }
    return num
}