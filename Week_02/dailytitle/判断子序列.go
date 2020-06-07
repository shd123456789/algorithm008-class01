// 时间复杂度O(n) 空间O(1) 2遍
func isSubsequence(s string, t string) bool {
    i,j := 0,0
    for j < len(s) && i < len(t) {
        if s[j] == t[i] {
            j++
        }
        i++
    }
    return j == len(s)
}

// 当s很多需要判断时预处理t通过二分查找将时间复杂度降低
func isSubsequence(s string, t string) bool {
    db := make([][]int, 26)
    for i:= 0; i < len(t); i++ {
       db[t[i] - 'a'] =  append(db[t[i] - 'a'], i)
    }
    
    prev := -1
    for i:= 0; i < len(s); i++ { 
        left := 0
        right :=  len(db[s[i] - 'a'])
        var mid int
        for left < right { // 改进方法通过二分查找到第一个大于prev的位置
            mid = left + (right - left) / 2
            if db[s[i] - 'a'][mid] > prev { // 如果大于prev继续循环直到left大于right时才能确定是第一个
                right = mid
            } else {
                left = mid + 1
            }
        }
        if left == len(db[s[i] - 'a']) {
            return false
        }
        prev = db[s[i] - 'a'][left]
    }
    return true
}

