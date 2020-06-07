// 时间复杂度O(n) 空间复杂度o(n)
func isValid(s string) bool {
    var stack []byte
    hash := map[byte]byte{'{':'}','[':']','(':')'}
    for i := 0; i < len(s); i++ {
        if s[i] == '{' || s[i] == '(' || s[i] == '['  {
            stack = append(stack, hash[s[i]])
        } else if len(stack) > 0 && stack[len(stack) - 1] == s[i] {
            stack = stack[:len(stack) - 1]
        } else {
            return false
        }
    }
    return len(stack) == 0
}
