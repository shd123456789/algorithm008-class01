// 栈
func removeOuterParentheses(S string) string {
    var stack,res []byte
    for i := 0; i < len(S); i++ {
        if S[i] == '(' {
            stack = append(stack, '(')
            if len(stack) > 1 {
                res = append(res, '(')
            }
        } else {
            stack = stack[:len(stack)-1]
            if len(stack) >= 1 {
                res = append(res, ')')
            }
        }
    }
    return string(res)
}

// 指针方法
func removeOuterParentheses(S string) string {
    k := 0
    res := make([]rune, 0)
    for _,num := range S {
        if num == ')' {
            k--
        }
        if k != 0 {
            res = append(res, num)
        }
        if num == '(' {
            k++
        }
    }
    return string(res)
}

