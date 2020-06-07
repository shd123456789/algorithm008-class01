// 时间复杂度为O（logn）空间为O(1)
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }
    l,r,mid := 0,len(matrix) * len(matrix[0]) - 1,0
    for l <= r {
        mid = l + (r - l) >> 1
        row := mid / len(matrix[0])
        col := mid % len(matrix[0])
        midV := matrix[row][col]
        if target == midV {
            return true
        } else if target > midV {
            l = mid + 1
        } else {
            r = mid - 1
        }
        
    }
    return false   
}