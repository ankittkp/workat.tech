package quickSort

func partition(start int, end int, arr []int) int {
	i := start + 1
	piv := arr[start] //make the first element as pivot element.
	for j := start + 1; j <= end; j++ {
		if arr[j] < piv {
			arr[i], arr[j] = arr[j], arr[i]
			i += 1
		}
	}
	arr[start], arr[i-1] = arr[i-1], arr[start] //put the pivot element in its proper place.
	return i - 1
}
func quickSort(low int, high int, arr []int) {
	if low < high {
		key := partition(low, high, arr)
		quickSort(low, key-1, arr)
		quickSort(key+1, high, arr)
	}
}
func QuickSort(arr []int) []int {
	quickSort(0, len(arr)-1, arr)
	return arr
}
