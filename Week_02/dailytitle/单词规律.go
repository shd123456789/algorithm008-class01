
// 时间复杂度O(n) 空间复杂度O(n) 2遍
func wordPattern(pattern string, str string) bool {
    strArr := strings.Split(str," ")
    if len(pattern) != len(strArr) {
        return false
    }
    hash,hash1 := make(map[byte]string), make(map[string]byte)
    for i := 0; i < len(pattern); i++ {
        if _,ok := hash[pattern[i]]; !ok {
            hash[pattern[i]] = strArr[i]
        } 
        if _,ok := hash1[strArr[i]]; !ok {
            hash1[strArr[i]] = pattern[i]
        }
        if hash[pattern[i]] != strArr[i] ||  hash1[strArr[i]] != pattern[i]  {
            return false
        }
    }
    return true
}