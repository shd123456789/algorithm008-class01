//11. 盛最多水的容器
func maxArea(height []int) int {
    len := len(height)
    left,right := 0, len - 1 
    maxArea := 0
    for i := 0; i < len; i++ {
        h1,h2 := height[left], height[right]
        maxArea = max(min(h1, h2) * (right - left), maxArea)
        if h1 > h2 {
            right--
        } else {
            left++
        }
    }
    return  maxArea
}

func min(n1 int,n2 int) int {
    if n1 > n2 {
        return n2
    } else {
        return n1
    }
}

func max(n1 int,n2 int) int {
    if n1 > n2 {
        return n1
    } else {
        return n2
    }
}