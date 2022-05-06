package mergeSort

/*
Implement Merge Sort
Given an array, sort it using merge sort.
5 4 2 5 3 1
3 11 4 200
Expected Output
1 2 3 4 5
4 11 200
*/

func merge(arr []int, l int, m int, r int) {
	var i, j, k int
	n1 := m - l + 1
	n2 := r - m

	L := make([]int, n1)
	R := make([]int, n2)

	for i = 0; i < n1; i++ {
		L[i] = arr[l+i]
	}
	for j = 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}
	i, j, k = 0, 0, l

	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}
	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}
func mergeSort(arr []int, low int, high int) []int {
	if low < high {
		mid := low + (high-low)/2

		mergeSort(arr, low, mid)
		mergeSort(arr, mid+1, high)

		merge(arr, low, mid, high)
	}
	return arr
}
func MergeSort(arr []int) []int {
	return mergeSort(arr, 0, len(arr)-1)
}
