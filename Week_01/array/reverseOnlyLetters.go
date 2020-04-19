// 917. 仅仅反转字母
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