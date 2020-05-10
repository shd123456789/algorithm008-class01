// 1遍
func topKFrequent(nums []int, k int) []int {
    hashFreg := make(map[int]int)
    for _,num := range nums {
        hashFreg[num]++
    }
    p := &priorityQueue{}
    heap.Init(p)
    for val,cnt := range hashFreg {
        heap.Push(p, element{val: val, freg:cnt})
        if p.Len() > k {
            heap.Pop(p)
        }
    }
    res := make([]int, 0) 
    for k > 0 {
        res = append(res, heap.Pop(p).(element).val)
        k--
    }
    return res
}

type element struct {
    val int
    freg int
}

type priorityQueue []element

func (q priorityQueue) Len() int           { return len(q) }
func (q priorityQueue) Less(i, j int) bool { return q[i].freg < q[j].freg }
func (q priorityQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *priorityQueue) Push(x interface{}) {
	*q = append(*q, x.(element))
}

func (q *priorityQueue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}