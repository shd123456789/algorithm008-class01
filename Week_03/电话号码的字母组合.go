var hash map[byte]string
var res []string
var path []byte
func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }
    hash = map[byte]string {
        2:"abc",
        3:"def",
        4:"ghi",
        5:"jkl",
        6:"mno",
        7:"pqrs",
        8:"tuv",
        9:"wxyz",
    }     
    res,path = make([]string, 0),make([]byte, 0)
    backtrack(digits, 0)
    return res
}

func backtrack(digits string, level int) {
    if len(path) == len(digits) {
        res = append(res, string(path))
        return
    }
    val := hash[digits[level] - '0']    
    for i := 0; i < len(val); i++ {
        path = append(path, val[i])
        backtrack(digits, level + 1)
        path = path[:len(path) - 1]
    }
}