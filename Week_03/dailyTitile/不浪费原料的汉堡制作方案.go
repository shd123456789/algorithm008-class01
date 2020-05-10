func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
    // 4x + 2y = tomatoSlices   x = (tomatoSlices - 2cheeseSlices)/2
    // 2x + 2y = 2cheeseSlices  y = cheeseSlices - x
    // 由以上可得只要保证x，y为大于等于0的整数，就由解
    
    c := tomatoSlices - 2 * cheeseSlices// 首先保证除数可以被2整除
    if c % 2  == 0 {
        x := c/2
        y := cheeseSlices - x
        if x >= 0 && y >=0 {
            return []int{x, y}
        }
    } 
    return []int{}
}