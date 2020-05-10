func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return nil
    }
    r,l,u,d := len(matrix[0]) - 1, 0, 0, len(matrix) - 1
    var res []int
    for true {
        for i := l; i <= r; i++ { // 向右遍历
            res = append(res, matrix[u][i])
        } 
        if u++; u > d { // 判断是否超过下边界
            break
        }
        for i := u; i <= d; i++ { // 向下遍历
            res = append(res, matrix[i][r])
        } 
        if r--; r < l { // 判断是否超过右边界
            break
        }
        for i := r; i >= l; i-- { // 向左遍历
            res = append(res, matrix[d][i])
        } 
        if d--; d < u { // 判断是否超过上边界
            break
        }
        for i := d; i >= u; i-- { // 向上遍历
            res = append(res, matrix[i][l])
        } 
        if l++; l > r { // 判断是否超过右边界
            break
        }
    }
    return res
}