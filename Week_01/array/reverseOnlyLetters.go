// 917. 仅仅反转字母 （2遍） 时间复杂度为O(n),空间复杂度为O(n) （3遍）
func reverseOnlyLetters(S string) string {
    
    f := 0; // 双指针
    l := len(S) - 1;
    buf := []byte(S)
    for f < l {
        if isLetter(buf[f]) == false {
            f++;
            continue
        }
        if isLetter(buf[l]) == false {
            l--;
            continue
        }
        buf[l], buf[f] = buf[f], buf[l]
        f++;
        l--;
    }

    return string(buf)
}

func isLetter(b byte) bool {
  return b >= 'A' && b <= 'Z' || b >= 'a' && b <= 'z'
}



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