func numDecodings(s string) int {
    if len(s) == 0 || s[0] == '0' {
        return 0
    }
    first,second := 1,0
    if s[0] != '0' {
        second = 1
    } 
    for i := 1; i < len(s); i++ {
        one := s[i] - '0'
        two := (s[i - 1] - '0') * 10 + s[i] - '0'
        ans := 0
        if 1 <= one && one <= 9 {
            ans += second
        }
        if 10 <= two && two <= 26 {
            ans += first
        } 
        first,second = second,ans
    }
    return second
}