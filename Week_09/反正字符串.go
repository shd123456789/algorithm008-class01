func reverseStr(s string, k int) string {
    sArr := []byte(s)
    for start := 0; start < len(sArr); start += 2 * k {
        l := min(start + k - 1, len(sArr) - 1)
        reverse(sArr, start, l)
    }
    return string(sArr)
}
// abcd.   2 
func reverse(s []byte, i, j int) {
    for i < j {
        s[i],s[j] = s[j],s[i]
        i++
        j--
    }
}

func min(i, j int) int {
    if i > j {
        return j
    } else {
        return i
    }
}