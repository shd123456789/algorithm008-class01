type Leaderboard struct {
    players [10001]int
    scores [10001]int
}


func Constructor() Leaderboard {
    return Leaderboard{}
}


func (this *Leaderboard) AddScore(playerId int, score int)  {
    os := this.players[playerId] 
    cs := this.players[playerId] + score
    this.players[playerId] = cs
    this.scores[cs] += 1
    if this.scores[os] > 0 {
        this.scores[os] -= 1
    } 
}


func (this *Leaderboard) Top(K int) int {
    res := 0
    for j := 10000; j >= 0; j-- {
        if l := this.scores[j] - K; l > 0 {
            res += K * j 
            break
        } else {
            K = -l
            res += this.scores[j] * j
        }
    }   
    return res
}


func (this *Leaderboard) Reset(playerId int)  {
    s := this.players[playerId] 
    if s > 0 {
        this.players[playerId] = 0
        this.scores[s] -= 1
        fmt.Println(s)
    }
}


/**
 * Your Leaderboard object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddScore(playerId,score);
 * param_2 := obj.Top(K);
 * obj.Reset(playerId);
 */