func quickSort(arr []int, start int, end int) {
	if end <= start {
		return
	}
	pivot := partition(arr, start, end)
	quickSort(arr, start, pivot - 1)
	quickSort(arr, pivot + 1, end)
}

func partition(arr []int, start int, end int) int {
	pivot,counter := end,start
	for i := start; i <= end; i++ {
		if arr[i] > arr[pivot] {
			arr[i],arr[counter] = arr[counter],arr[i]
			counter++
		}
	}
	arr[counter],arr[pivot] = arr[pivot],arr[counter]
	return counter
}