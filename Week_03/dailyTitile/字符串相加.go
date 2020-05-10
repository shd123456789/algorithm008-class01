func addStrings(num1 string, num2 string) string {
    l,l1,l2 := 0,len(num1),len(num2) 
    var res []byte
    var carr byte
    if l1 > l2 {
        l,res,l1,l2 = l1,make([]byte, l1 + 1),l1-1,l2-1
    } else {
        l,res,l1,l2 = l2,make([]byte, l2 + 1),l1-1,l2-1
    } 
    for l1 >= 0 || l2 >= 0 || carr > 0 {
        n := carr
        if l1 >=0 {
           n += num1[l1] - '0'
           l1--
        }
        if l2 >=0 {
           n += num2[l2] - '0'
           l2--
        }
        carr = n / 10
        res[l] = (n % 10) + '0'
        l--
    } 
    if res[0] == byte(0) {
        res = res[1:len(res)]
    }
    return string(res)
}