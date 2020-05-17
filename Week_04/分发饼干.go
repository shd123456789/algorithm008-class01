//时间复杂度为O(N) 空间为(1) 2遍
func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    x,y := 0,0
    for x < len(g) && y < len(s) {
        if s[y] >= g[x] {
            x++
        }
        y++
    }
    return x
}