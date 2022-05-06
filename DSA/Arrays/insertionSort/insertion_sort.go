package insertionSort

/*
Given an array, sort it using insertion sort.
Sample Input
2
5 4 2 3 1

3 11 4 200
Expected Output
1 2 3 4 5
3 4 11 200
*/

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = key
	}
	return arr
}
