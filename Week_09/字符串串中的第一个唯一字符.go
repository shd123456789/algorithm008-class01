func firstUniqChar(s string) int {
    var hash [26]int
    for i := 0; i < len(s); i++ {
        hash[s[i] - 'a']++
    }
    for i := 0; i < len(s); i++ {
        if hash[s[i] - 'a'] == 1 {
            return i
        }
    }
    return -1
}