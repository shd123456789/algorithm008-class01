// 88. 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int)  {
    sl := m + n - 1
    l1 := m - 1
    l2 := n - 1
    for l1 >= 0 && l2 >= 0 {
        if nums1[l1] > nums2[l2] {
            nums1[sl] = nums1[l1]
            l1--
        } else {
            nums1[sl] = nums2[l2]
            l2--
        }
        sl--
    } 
    if l2 >= 0 {
        for i := 0; i <= sl; i++ {
            nums1[i] = nums2[i]
        }
    }
}