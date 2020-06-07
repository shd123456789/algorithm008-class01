var minCount int
var stepRecord map[string]bool 
func minMutation(start string, end string, bank []string) int {
    minCount = len(start) + 1
    stepRecord = make(map[string]bool)
    dfs(start, end, 0, bank)
    if minCount == len(start) + 1 {
        minCount = -1
    }
    return minCount
}

func dfs(str string, end string, stepCount int, bank []string) {
    if str == end || stepCount > minCount {
        if stepCount < minCount {
            minCount = stepCount
        }
        return 
    }
    for _,v := range bank {
        if stepRecord[v] {
            continue
        }
        diff := 0
        for i := 0; i < len(v); i++ {
            if v[i] != str[i] {
                if diff++; diff > 1 {
                    break
                }
            }
        }
        if diff == 1 {
            stepRecord[v] = true
            dfs(v, end, stepCount + 1, bank)
            delete(stepRecord, v)
        }
    }
}

