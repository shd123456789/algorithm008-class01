// 时间O(n) 除结果集外空间O(1) 
func replaceSpace(s string) string {
    var res []byte
    for i := 0; i < len(s); i++ {
        if s[i] == ' ' {
            res = append(res, '%', '2', '0')
        } else {
            res = append(res, s[i])
        }
    }
    return string(res)
}