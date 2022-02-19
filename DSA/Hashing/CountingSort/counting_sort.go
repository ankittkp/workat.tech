package CountingSort

import (
	"math"
)

func CountingSort(arr []int) []int{
	max, min:= math.MinInt, math.MaxInt
	for _, value := range arr {
		if value > max{
			max = value
		}
		if value <= min{
			min = value
		}
	}

	count := make([]int, max-min+1)
	output := make([]int, len(arr))
	for index :=0;index<len(arr);index++{
		count[arr[index]-min]++
	}
	for index :=1;index<max-min+1;index++{
		count[index] += count[index-1]
	}
	for i:= 0; i <len(arr); i++ {
		output[count[arr[i]-min] - 1] = arr[i]
		count[arr[i]-min]--
	}
	for index :=0;index<len(arr);index++{
		arr[index] = output[index]
	}
	return arr

}