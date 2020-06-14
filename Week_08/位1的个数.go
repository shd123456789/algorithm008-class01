func hammingWeight(num uint32) int {
    s := 0
    for {
        if num == 0 {
            break
        }
        num = num&(num - 1)
        s++
    }
    return s
}