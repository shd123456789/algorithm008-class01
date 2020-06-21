func myAtoi(str string) int {
    index,sign,total := 0,1,0
    if len(str) == 0 {
        return 0
    }
    for index < len(str)  &&  str[index] == ' ' {
        index++
    }
    if index == len(str) {
        return 0
    }
    if str[index] == '-' || str[index] == '+' {
        if str[index] == '-' {
            sign = -1
        }
        index++
    }
    for index < len(str) {
        dig := int(str[index] - '0')
        if dig > 9 || dig < 0 {
            break
        }
        if (math.MaxInt32 / 10) < total || 
            ((math.MaxInt32 / 10) == total && (math.MaxInt32 % 10) < dig) {
                if sign == -1 {
                    return math.MinInt32
                } else {
                    return math.MaxInt32
                }
            break
        }
        total = total * 10 + dig
        index++
    }
    return total * sign
}