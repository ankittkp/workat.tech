package identicalTwins

/*
For an array of integers nums, an identical twin is defined as pair (i, j) where nums[i] is equal to nums[j] and i < j.

Array: [1, 2, 3, 2, 1]
Number of Identical Twins: 2
Explanation:
Identical Twins: [[1, 1], [2, 2]]
Indexes: (0, 4), (1, 3)
Array: [1, 2, 2, 3, 2, 1]
Number of Identical Twins: 4
Explanation:
Identical Twins: [[1, 1], [2, 2], [2, 2], [2, 2]]
Indexes: (0, 5), (1, 2), (1, 4), (2, 4)
Array: [1, 1, 1, 1]
Number of Identical Twins: 6
Explanation:
Identical Twins: [[1, 1], [1, 1], [1, 1], [1, 1], [1, 1], [1, 1]]
Indexes: (0, 1), (0, 2), (0, 3), (1, 2), (1, 3), (2, 3)
Given an array nums, find the number of identical twins.

*/

const LIMIT int = 10000

func identicalTwins(arr []int) int {
	s := make([]int, LIMIT)
	for i := 0; i < len(arr); i++ {
		s[arr[i]]++
	}
	sum := 0
	for i := 0; i < len(s); i++ {
		if s[i] > 0 {
			sum += s[i] * (s[i] - 1) / 2
		}
	}
	return sum
}
