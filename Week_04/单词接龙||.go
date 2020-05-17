func findLadders(beginWord string, endWord string, wordList []string) [][]string {
    res := [][]string{}
    flag := false
    for i := 0; i < len(wordList); i++ {
        if wordList[i] == endWord {
            flag = true
        }
    }
    if !flag {
        return res
    }
    preIndex := make(map[string][]string)
    for _,word := range wordList {
        for i := 0; i < len(word); i++ {
            key := word[:i] + "*" + word[i+1:]
            preIndex[key] = append(preIndex[key], word)
        }
    }
    queue := [][]string{{beginWord}}
    visited := make(map[string]int)
    visited[beginWord] = -1
    flag = false
    level := 0
    for len(queue) > 0 && !flag {
        level++
        l := len(queue)
        for i := l - 1; i >= 0; i-- {
            path := queue[i]
            w := path[len(path) - 1]
            for j := 0; j < len(w); j++ {
                key := w[:j] + "*" + w[j+1:]
                set := preIndex[key]
                for _,v := range set {
                    if visited[v] == 0 || visited[v] == level {
                        visited[v] = level
                        newPath := append([]string{}, path...)
                        newPath = append(newPath, v)
                        queue = append(queue, newPath)
                        if v == endWord {
                            flag= true
                            res = append(res, newPath)
                        }
                    }
                }
            }
        }
        queue = queue[l:]
    }
    return res
}