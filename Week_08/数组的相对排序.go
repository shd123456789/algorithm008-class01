func relativeSortArray(arr1 []int, arr2 []int) []int {
    hash := make([]int, 1001)
    res := []int{}
    for _,v := range arr1 {
        hash[v]++
    }
    for _,v := range arr2 {
        for i := 0; i < hash[v]; i++ {
            res = append(res, v)
        }
        hash[v] = 0
    }
    tmp := []int{}
    for _,v := range arr1 {
        if hash[v] != 0 {
            tmp = append(tmp, v)
        }
    }
    sort.Ints(tmp)
    res = append(res, tmp...)
    return res
}