func ladderLength(beginWord string, endWord string, wordList []string) int {
    set := make(map[string]bool, 0)
	for _, word := range wordList {
		set[word] = true
	}
    if !set[endWord] {
        return 0
    }
    visited,endVisited := make(map[string]int),make(map[string]int)
    preIndex := make(map[string][]string)
    queue,endQueue := []string{beginWord},[]string{endWord}
    visited[beginWord],endVisited[endWord] = 1,1
    for _,word := range wordList {
        for i := 0; i < len(word); i++ {
            s := word[:i] + "*" + word[i+1:]
            preIndex[s] = append(preIndex[s], word)
        } 
    }
	for len(queue) > 0 && len(endQueue) > 0 {
        if len(queue) > len(endQueue) {
            visited,endVisited = endVisited,visited
            queue,endQueue = endQueue,queue
        }
        l := len(queue)
        level := visited[queue[0]]
        for i := 0; i < l; i++ {
            w := queue[i]
            for j := 0; j < len(w); j++ {
                k := w[:j] + "*" + w[j+1:]
                set := preIndex[k]
                for l := 0; l < len(set); l ++ {
                    if _,ok := endVisited[set[l]]; ok {
                        return level + endVisited[set[l]]
                    }
                    if _,ok := visited[set[l]]; !ok {
                        visited[set[l]] = level + 1
                        queue = append(queue, set[l])
                    }
                }
            }
        }
        queue = queue[l:]
    }
    return 0

}