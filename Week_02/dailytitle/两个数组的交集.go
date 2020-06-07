// 简单方法 3遍
func intersect(nums1 []int, nums2 []int) []int {
    hash := make(map[int]int)
    var res []int
    for i := 0; i < len(nums1); i++ {
        hash[nums1[i]]++
    }
    for i := 0; i < len(nums2); i++ {
        if hash[nums2[i]]--; hash[nums2[i]] >= 0 {
            res = append(res, nums2[i])
        } 
    }
    return res
}

// 进阶1
func intersect(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    i,j := 0,0
    var res []int
    for i < len(nums1) && j < len(nums2) {
        if nums1[i] == nums2[j] {
            res = append(res, nums2[j])
            i++
            j++
        } else if nums1[i] > nums2[j] {
            j++
        } else {
            i++
        }
    }
    return res
}
