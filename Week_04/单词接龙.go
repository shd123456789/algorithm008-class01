// 时间复杂度O(M*N) M为单词的长度 N为单词表的总数 空间复杂度O(M*N) 2遍
func ladderLength(beginWord string, endWord string, wordList []string) int {
    preHash := make(map[string][]string) //预处理
    for _,w := range wordList {
        for i := 0; i < len(w); i++ {
            k := w[:i] + "*" + w[i+1:]
            preHash[k] = append(preHash[k], w)
        }
    }
    queue := []string{beginWord}
    visited := make(map[string]bool)
    visited[beginWord] = true
    count := 1
    for len(queue) > 0 {
        count++
        l := len(queue)
        for i := 0; i < l; i++ {
            w := queue[i]
            for j := 0; j < len(w); j++ {
                k := w[:j] + "*" + w[j+1:]
                set := preHash[k]
                for _,e := range set {
                    if e == endWord {
                        return count
                    }
                    if !visited[e] {
                        visited[e] = true
                        queue = append(queue, e)
                    }
                }
            }
        }
        queue = queue[l:]
    }
    return 0
}

