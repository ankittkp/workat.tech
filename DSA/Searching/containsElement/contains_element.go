package containsElement

/*
Given a sorted array and a number key, return whether the key is present in the array or not.

Expected Time Complexity: O(log n)

Examples
Array: [1, 2, 3, 3, 3, 4, 4, 5]
Number: 2
Answer: true
Array: [1, 2, 3, 3, 3, 4, 4, 5]
Number: 6
Answer: false
*/

func BinarySearch(low int, high int, arr []int, num int) bool {
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == num {
			return true
		} else if arr[mid] > num {
			return BinarySearch(low, mid-1, arr, num)
		} else {
			return BinarySearch(mid+1, high, arr, num)
		}
	}
	return false
}
func ContainsElement(arr []int, num int) bool {
	low, high := 0, len(arr)
	return BinarySearch(low, high-1, arr, num)
}
