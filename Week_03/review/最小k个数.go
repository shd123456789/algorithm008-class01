// 1遍
func getLeastNumbers(arr []int, k int) []int {
    var h IntHeap
    heap.Init(&h)
    var res []int
    for i := 0; i < len(arr); i++ {
        heap.Push(&h, arr[i])
    }
    for len(res) < k {
        res = append(res, heap.Pop(&h).(int))
    }
    return res
}

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
  res := (*h)[len(*h)-1]
  *h = (*h)[:len(*h)-1]
  return res
}