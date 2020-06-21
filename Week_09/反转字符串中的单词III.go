func reverseWords(s string) string {
    p,i := 0,0
    sArr := []byte(s)
    for i < len(sArr) {
        if s[i] == ' ' {
            reverse(sArr, p, i - 1)
            p = i+1
        }
        i++
    }
    reverse(sArr, p, len(sArr) - 1)
    return string(sArr)
}

func reverse(s []byte, i, j int) {
    for i < j {
        s[i],s[j] = s[j],s[i]
        i++
        j--
    }
}