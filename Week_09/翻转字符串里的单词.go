func reverseWords(s string) string {
    s = strings.Trim(s, " ")
    start,l := len(s) - 1,0
    res := []byte{}
    for start >= 0 {
        if s[start] != ' ' {
            l++
            start--
        } else {
            res = append(res, s[start + 1 : start + 1 + l]...)
            res = append(res, ' ')
            start--
            for start >= 0 && s[start] == ' ' {
                start--
            }
            l = 0
        }
    } 
    if l > 0 {
        res = append(res, s[:l]...)
    }
    return string(res)
}