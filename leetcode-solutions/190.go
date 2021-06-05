func reverseBits(num uint32) uint32 {
    out := uint32(0)
    k := uint32(31)
    for num > 0 {
        out |= (num&1) << k
        num >>= 1
        k--
    }
    return out
}
