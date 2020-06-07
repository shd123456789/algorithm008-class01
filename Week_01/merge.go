// 88. 合并两个有序数组(时间复杂度为O(n)空间为O(1)) 2遍
func merge(nums1 []int, m int, nums2 []int, n int)  {
    s,l1,l2 := m + n - 1,m - 1,n - 1
    for l1 >= 0 && l2 >= 0 {
        if nums1[l1] > nums2[l2] {
            nums1[s] = nums1[l1]
            l1--
        } else {
            nums1[s] = nums2[l2]
            l2--
        }
        s--
    }
    if l2 >= 0 {
        for i := 0; i <= l2; i++ {
            nums1[i] = nums2[i]
        }
    }
}