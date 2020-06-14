func reverseBits(num uint32) uint32 {
    res := uint32(0)
    t := 32
    for t != 0 {
        res <<= 1
        res |= (num & 1)
        num >>= 1
        t--
    }
    return res
}