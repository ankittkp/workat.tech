package positiveCumulativeSum

/*
The cumulative sum of an array at index i is defined as the sum of all elements of the array from index 0 to index i.

Examples
Initial Array: [1, -2, 3, 4, -6]
Cumulative Sum: [1, -1, 2, 6, 0]
Initial Array: [1, -1, -1, -1, 1]
Cumulative Sum: [1, 0, -1, -2, -1]
Initial Array: [1, 3, 5, 7]
Cumulative Sum: [1, 4, 9, 16]
The positive cumulative sum of an array is a list of only those cumulative sums which are positive.

Initial Array: [1, -2, 3, 4, -6]
Cumulative Sum: [1, -1, 2, 6, 0]
Positive Cumulative Sum: [1, 2, 6]
Initial Array: [1, -1, -1, -1, 1]
Cumulative Sum: [1, 0, -1, -2, -1]
Positive Cumulative Sum: [1]

*/

func positiveCumSum(arr []int) []int {
	j := 0
	res := []int{arr[0]}
	for i := 1; i < len(arr); i++ {
		arr[i] += arr[i-1]
		if arr[i] > 0 {
			res = append(res, arr[i])
			j++
		}
	}
	return res
}
