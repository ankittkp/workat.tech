package cumulative_sum

/*
The cumulative sum of an array at index i is defined as the sum of all elements of the array from index 0 to index i.

Examples
Initial Array: [1, 2, 3, 4]
Cumulative Sum: [1, 3, 6, 10]

Initial Array: [1, 1, 1, 1, 1]
Cumulative Sum: [1, 2, 3, 4, 5]

Initial Array: [1, 3, 5, 7, 9]
Cumulative Sum: [1, 4, 9, 16, 25]
 */

func CumSum(arr [] int) [] int{
	for i:=1; i<len(arr);i++ {
		arr[i] += arr[i-1]
	}
	return arr
}
