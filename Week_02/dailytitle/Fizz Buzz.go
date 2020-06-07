func fizzBuzz(n int) []string {
    var res []string 
    for i := 1; i <= n; i++ {
        if i % 3 == 0 && i % 5 == 0 {
           res = append(res, "FizzBuzz") 
        } else if i % 5 == 0 {
            res = append(res, "Buzz")
        } else if i % 3 == 0 {
            res = append(res, "Fizz")
        } else {
            res = append(res, strconv.Itoa(i))
        }
    }
    return res
}
   2 
[1,2,3,4]

n - i,
n - i + 1

1
2
4
8
16