package IdenticalTwins

/*
For an array of integers nums, an identical twin is defined as pair (i, j) where nums[i] is equal to nums[j] and i < j.

Examples
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

type data struct {
	count int
	index []int
}

func identicalTwins(arr []int) int {
	m := make(map[int]data)
	for i, value := range arr {
		m[value] = data{
			count: m[value].count + 1,
			index: append(m[value].index, i),
		}
	}
	output := 0
	for _, value := range m {
		output += (value.count * (value.count - 1)) / 2
	}
	return output

}
