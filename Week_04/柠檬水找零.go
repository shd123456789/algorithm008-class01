// 时间复杂度为O(n) 空间复杂度为O(1) 2遍
func lemonadeChange(bills []int) bool {
   fiveNum := 0
   tenNum := 0
   for i := 0; i < len(bills); i++ {
       if bills[i] == 5 {
           fiveNum++
       } else if bills[i] == 10 {
           if fiveNum == 0 {
               return false
           }
           fiveNum--
           tenNum++
       } else {
           if tenNum > 0 && fiveNum > 0 {
               tenNum--
               fiveNum--
           } else if fiveNum >= 3 {
               fiveNum -= 3
           } else {
               return false
           }
       }
   }
   return true
}