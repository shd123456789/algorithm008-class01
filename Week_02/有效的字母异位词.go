// 时间O(N) 空间O(1)
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    arr := make([]int, 26)
    for i := 0;i < len(s); i++ {
        arr[s[i] - 'a']++ 
    }
    for i := 0;i < len(t); i++ {
        if arr[t[i] - 'a']--; arr[t[i] - 'a'] < 0 {
            return false
        } 
    }
    return true
}