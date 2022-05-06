package evenNoOfDigits

/*
Given an array of integers, find the elements which have an even number of digits.

Example
Array: [42, 564, 5775, 34, 123, 454, 1, 5, 45, 3556, 23442]
Answer: 42, 5775, 34, 45, 3556
The order of the returned elements should be the same as the order of the initial array.

*/
func countDigits(x int) int {
	count := 0
	for x != 0 {
		count++
		x = x / 10
	}
	return count
}
func evenNoOfDigits(arr []int) []int {
	var res []int
	for i := 0; i < len(arr); i++ {
		if got := countDigits(arr[i]); got%2 == 0 && got != 0 {
			res = append(res, arr[i])
		}
	}
	return res
}
