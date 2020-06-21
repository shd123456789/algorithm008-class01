func findAnagrams(s string, p string) []int {
    window,needed,valid := make(map[byte]int),make(map[byte]int),0
    for i := 0; i < len(p); i++ {
        needed[p[i]]++
    }
    res,left,right := []int{},0,0
    for right < len(s){
        c := s[right]
        right++
        if _,ok := needed[c]; ok {
            window[c]++
            if window[c] == needed[c] {
                valid++
            }
        } 
        for right - left >= len(p) {
            if valid == len(needed) {
                res = append(res, left)
            }
            c1 := s[left]
            left++
            if _,ok := needed[c1]; ok {
                if window[c1] == needed[c1] {
                    valid--
                }
                window[c1]--
            }
        }
    }
    return res
}