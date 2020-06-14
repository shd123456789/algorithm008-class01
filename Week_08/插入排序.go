func insertSort(arr []int) {
	prev,current := 0,0
	for i := 1;i < len(arr); i++ {
		prev = i - 1
		current = arr[i]
		for prev >= 0 && arr[prev] > current {
			arr[prev + 1] = arr[prev]
			prev--
		}
		arr[prev + 1] = current
	}
}