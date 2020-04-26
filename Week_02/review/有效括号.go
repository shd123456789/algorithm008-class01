// 有效括号.go
func isValid(s string) bool {
    var stack []byte
    keyMap := map[byte]byte{'[': ']','{': '}','(': ')',}
    for i := 0; i < len(s); i++ {
        if s[i] == '(' || s[i] == '{' || s[i] == '[' {
            stack = append(stack, s[i])
        } else if len(stack) > 0 && s[i] == keyMap[stack[len(stack) - 1]] {
            stack = stack[:len(stack)-1]
        } else { // 当碰到只有右括号时， 立马跳出循环
            return false
            break
        }
    } 
    return len(stack) == 0
}