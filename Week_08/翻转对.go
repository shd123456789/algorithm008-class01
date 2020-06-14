var count int
func reversePairs(nums []int) int {
    count = 0
    mergeSort(nums, 0, len(nums) - 1)
    return count
}

func mergeSort(arr []int, left int, right int) {
	if right <= left {
		return
	}
	mid := left + (right - left) >> 1
	mergeSort(arr, left, mid)
	mergeSort(arr, mid + 1, right)
	merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) {
	tmp := make([]int, right - left + 1)
	l,r,k := left, mid + 1, 0

    for i,j := l,mid + 1; i <= mid; i++ {
        for j <= right && (float64(arr[i]) / 2.00) > float64(arr[j]) {
            j++
        }
        count += j - mid - 1
    }  

	for l <= mid && r <= right {
		if arr[l] < arr[r] {
			tmp[k] = arr[l]
			l++
		} else {
			tmp[k] = arr[r]
			r++
		}
		k++
	}
	for l <= mid  {
		tmp[k] = arr[l]
		l++
		k++
	}
	for r <= right  {
		tmp[k] = arr[r]
		r++
		k++
	}
	for i := 0; i < len(tmp); i++ {
		arr[left + i] = tmp[i]
	}
}