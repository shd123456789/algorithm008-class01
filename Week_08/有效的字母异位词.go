func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    var hash [26]int
    for i := 0; i < len(s); i++ {
        hash[s[i] - 'a']++
    }
    for i := 0; i < len(t); i++ {
        if hash[t[i] - 'a']--; hash[t[i] - 'a'] < 0 {
            return false
        }
    }
    return true
}