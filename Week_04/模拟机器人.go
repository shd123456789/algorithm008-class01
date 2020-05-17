// 2遍
func robotSim(commands []int, obstacles [][]int) int {
    hashObstacles := make(map[[2]int]bool)
    for _,pos := range obstacles {
        hashObstacles[[2]int{pos[0],pos[1]}] = true
    }
    x,y,d := 0,0,1
    maxDistance := math.MinInt64
    dPos := [][]int{{0,1},{1,0},{0,-1},{-1,0}}
    for i := 0; i < len(commands); i++ {
        if commands[i] == -1 {
            if d++; d > 4 {
                d = 1
            }
        } else if commands[i] == -2 {
            if d--; d < 1 {
                d = 4
            }
        } else {
            for j := 0; j < commands[i]; j++ {
                tmpX,tmpY := x + dPos[d - 1][0],y + dPos[d - 1][1]
                if hashObstacles[[2]int{tmpX, tmpY}] {
                    break
                }
                x,y = tmpX,tmpY
            }
            distance := x * x + y * y
            if maxDistance < distance {
                maxDistance = distance
            }
        }

    }
    return maxDistance
}

