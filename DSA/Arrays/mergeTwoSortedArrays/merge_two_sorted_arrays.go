package mergeTwoSortedArrays

/*
Given two sorted arrays A and B, find the merged sorted array C by merging A and B.

Example:
A: [1, 2, 3, 4, 4]
B: [2, 4, 5, 5]
C: [1, 2, 2, 3, 4, 4, 4, 5, 5]
*/

func MergeTwoSortedArrays(firstArray []int, secondArray []int) []int {
	res := make([]int, len(firstArray)+len(secondArray))
	i, j, k := 0, 0, 0
	for i < len(firstArray) && j < len(secondArray) {
		if firstArray[i] < secondArray[j] {
			res[k] = firstArray[i]
			i++
		} else {
			res[k] = secondArray[j]
			j++
		}
		k++
	}
	for i < len(firstArray) {
		res[k] = firstArray[i]
		i++
		k++

	}
	for j < len(secondArray) {
		res[k] = secondArray[j]
		j++
		k++
	}
	return res
}
