func addDigits(num int) int {
    tmp := 0
    for num != 0 {
       tmp += num %  10
       if num /= 10; num == 0 {
           num,tmp = tmp,0
           if num < 10 {
               break
           }
       }
    }
    return num
}   