// 字母异味词分组 时间复杂度O(M*N)N为数组长度M为字符串最大长度 2遍
func groupAnagrams(strs []string) [][]string {
    var res [][]string
    if len(strs) == 0 {
        return res
    }
    hash := make(map[[26]byte][]string)
    for i := 0; i < len(strs); i++ {
        key := [26]byte{}
        for j := 0; j < len(strs[i]); j++ {
            key[strs[i][j] - 'a']++
        }
        hash[key] = append(hash[key], strs[i])
    }
    for _,v := range hash {
        res = append(res, v)
    }
    return res
}


