package mergeSortedSubarrays

import "weavelab.xyz/workat.tech/DSA/Arrays/mergeTwoSortedArrays"

/*
Consider an array that is divided into two parts and both of the parts are sorted individually.
Given the mentioned array and the end index of the first part, merge them to create a sorted array.
Update the same array with the merged elements. You can use extra auxiliary space.

Expected Time Complexity: O(n) where n denotes the size of the array.

Example
Array: [1, 3, 5, 7, 9, 11, 0, 4, 6, 8]
End Index: 5
Array after merging: [0, 1, 3, 4, 5, 6, 7, 8, 9, 11]
*/
func mergeSortedSubarrays(arr []int, endIndex int) []int {
	firstArray := arr[:endIndex+1]
	secondArray := arr[endIndex+1:]
	res := mergeTwoSortedArrays.MergeTwoSortedArrays(firstArray, secondArray)
	return res
}
