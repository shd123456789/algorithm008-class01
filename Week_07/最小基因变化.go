func minMutation(start string, end string, bank []string) int {
    visited := make(map[string]bool)
    preIndex := make(map[string][]string)
    queue := []string{start}
    visited[start] = true
    count := 0
    for _,word := range bank {
        for i := 0; i < len(word); i++ {
            s := word[:i] + "*" + word[i+1:]
            preIndex[s] = append(preIndex[s], word)
        } 
    }
	for len(queue) > 0 {
        count++
        l := len(queue)
        for i := 0; i < l; i++ {
            w := queue[i]
            for j := 0; j < len(w); j++ {
                k := w[:j] + "*" + w[j+1:]
                set := preIndex[k]
                for l := 0; l < len(set); l ++ {
                    if set[l] == end {
                        return count
                    }
                    if !visited[set[l]] {
                        visited[set[l]] = true
                        queue = append(queue, set[l])
                    }
                }
            }
        }
        queue = queue[l:]
    }
    return -1

}