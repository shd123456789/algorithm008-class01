func reverseOnlyLetters(S string) string {
    left,right := 0,len(S) - 1
    str := []byte(S)
    for left < right {
        if !isLetter(str[left]) {
            left++
            continue
        }
        if !isLetter(str[right]) {
            right--
            continue
        }
        str[left],str[right] = str[right],str[left]
        left++
        right--
    }
    return string(str)

}

func isLetter(s byte) bool {
    return 'A' <= s && s <= 'Z' || 'a' <= s && s <= 'z'
}