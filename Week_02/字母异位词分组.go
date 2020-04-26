func groupAnagrams(strs []string) [][]string {
    var res [][]string
    hashRes := make(map[[26]byte][]string)
    for i := 0; i < len(strs); i++ {
        var key [26]byte
        for j := 0; j < len(strs[i]); j++ {
            key[strs[i][j] - 'a']++   
        }
        if _,ok := hashRes[key]; ok {
            hashRes[key] = append(hashRes[key], strs[i])
        } else {
            hashRes[key] = []string{strs[i]}
        }
    } 
    for _,v := range hashRes {
        res = append(res, v)
    }
    return res
}