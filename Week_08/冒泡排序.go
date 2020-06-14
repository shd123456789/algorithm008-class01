func bubbleSort(arr []int) {
	for i := 0;i < len(arr); i++ {
		flag := false
		for j := 0;j < len(arr) - 1 - i; j++ {
			if arr[j] > arr[j+1] {
				arr[j],arr[j+1] = arr[j + 1],arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}