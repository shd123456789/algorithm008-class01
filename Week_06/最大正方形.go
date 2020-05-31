func maximalSquare(matrix [][]byte) int {
    maxLen := byte(0)
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[0]); j++ {
            if matrix[i][j] == '0' || i == 0 || j == 0 {
                if matrix[i][j] == '1' && maxLen == byte(0) {
                    maxLen = byte(1)
                }
                continue
            }
            mi := min(matrix[i - 1][j], matrix[i][j - 1])
            mi = min(mi, matrix[i - 1][j - 1])
            matrix[i][j] = mi + 1
            maxLen = max(maxLen, matrix[i][j] - '0')
        }  
    }
    r := int(maxLen)
    return r * r
}
func min (n1 byte, n2 byte) byte {
    if n1 > n2 {
        return n2
    } else {
        return n1
    }
}

func max (n1 byte, n2 byte) byte {
    if n1 > n2 {
        return n1
    } else {
        return n2
    }
}