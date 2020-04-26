func getHint(secret string, guess string) string {
    hash := make(map[byte]int)
    a := 0
    b := 0
    for i := 0; i < len(secret); i++  {
        if secret[i] == guess[i] {
            a++
        }
        hash[secret[i]]++
    } 
    for i := 0; i < len(guess); i++  {
        if _,ok := hash[guess[i]];ok {
            hash[guess[i]] = hash[guess[i]] - 1
            if hash[guess[i]] >= 0 {
                b++
            }
        }
    }
    b = b - a
    return strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
}

func getHint(secret string, guess string) string {
    arr := make([]int, 10)
    bulls := 0
    cows := 0
    for i := 0; i < len(secret); i++  {
        if b,c := secret[i] - '0', guess[i] - '0'; b == c {
            bulls++
        } else {
            if ((arr[b]) < 0) {
                cows++
            }
            if ((arr[c]) > 0) {
                cows++
            }
            arr[b] ++
            arr[c] --
        }
    } 
    return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}