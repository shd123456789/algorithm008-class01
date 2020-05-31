func leastInterval(tasks []byte, n int) int {
    s := make([]int, 26)
    for _,v := range tasks {
        s[v - 'A']++
    }
    sort.Ints(s)
    maxCount := s[len(s) - 1]
    resCount := (maxCount - 1) * (n + 1) + 1
    i := 24
    for i >= 0 && s[i] == maxCount {
        resCount++
        i--
    }
    if resCount < len(tasks) {
        resCount = len(tasks)
    }
    return resCount
}