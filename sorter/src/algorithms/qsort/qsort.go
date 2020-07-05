/*
# author: @monkeyrp
*/
package qsort

import "fmt"

func _QuickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right
	for i <= j {
		for j >= p && values[j] >= values[p]{
			j--
		}

		if j >= p {
			values[p], values[j] = values[j], values[p]
			p = j
		}

		if i <= p && values[i] <= temp{
			i++
		}
		if i <= p {
			values[p], values[i] = values[i], values[p]
			p = i
		}
		//values[p] = temp
		if p - left > 1 {
			_QuickSort(values, left, p-1)
		}
		if right - p > 1 {
			_QuickSort(values, p+1, right)
		}
	}
	
	fmt.Println("finally sort")
	for index, value := range values{
		fmt.Println("index: ", index, "value: ", value)
	}
}

func QuickSort(values []int) {
	_QuickSort(values, 0, len(values) - 1)
}