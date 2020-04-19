// 移动零
func moveZeroes(nums []int)  { // 1遍
    len := len(nums) - 1
    if len <= 0 {
        return
    }

    insertPose := 0;
    for _,num := range nums {
        if num != 0 {
            nums[insertPose] = num 
            insertPose ++
        }
    } 

    for {
        if len < insertPose {
            break
        }
        nums[insertPose] = 0;
        insertPose ++
    }
}