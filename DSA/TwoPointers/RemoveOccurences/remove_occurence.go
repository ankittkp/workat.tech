package Remove_Occurences

/*
Given an array and a number k, remove all occurrences of k from the array (in-place). Return the number of elements 'remainingSize' left after removing k.
There can be anything beyond the first 'remainingSize' elements. It will be ignored.

Example
Array: [1, 4, 2, 6, 2, 6, 9, 4]
Number: 4
Answer: 6
Explanation:[1, 2, 6, 2, 6, 9]
 */

func RemoveOccurences(arr [] int, k int) int {
	count := 0
	for _, value := range arr {
		if value == k {
			count++
		}
	}
	return len(arr) - count
}