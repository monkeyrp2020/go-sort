/*
# author: @monkeyrp
*/
package bubblesort

import "fmt"
import "testing"

func TestBubblesort1(t *testing.T) {
	values := []int {1, 67, 4, 234235, 9123, 1111111}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 4 || values[2] != 67 || values[3] != 9123 || values[4] != 234235 || values[5] !=  1111111 {
		fmt.Println("the BulleSort algorithms is false.")
		t.Error("BulleSort() failed. Got", values, "Expected 1 4 67 9123 234235 1111111")
	}
}

func TestBubblesort2(t *testing.T) {
	values := []int {1, 67, 4, 234235, 9123, 234235}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 4 || values[2] != 67 || values[3] != 9123 || values[4] != 234235 || values[5] !=  234235 {
		fmt.Println("the BulleSort algorithms is false.")
		t.Error("BulleSort() failed. Got", values, "Expected 1 4 67 9123 234235 234235")
	}
}

func TestBubblesort3(t *testing.T) {
	values := []int {234235}
	BubbleSort(values)
	if values[0] !=  234235 {
		fmt.Println("the BulleSort algorithms is false.")
		t.Error("BulleSort() failed. Got", values, "Expected 234235")
	}
}