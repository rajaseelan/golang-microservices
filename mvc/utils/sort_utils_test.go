package utils

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	els := []int{9, 8, 7, 6, 5}

	els = BubbleSort(els)

	if els == nil {
		t.Error("Not expecting els to be nil")
	}

	if len(els) != 5 {
		t.Error("Expected 5 elements in return")
	}

	if els[0] != 5 {
		t.Error("Expected 5 for element")
	}

	if els[2] != 7 {
		t.Error("Expected 7 for element")
	}

	if els[4] != 9 {
		t.Error("Expected 9 for element")
	}

}

func TestBubbleSortNilSlice(t *testing.T) {

	BubbleSort(nil)

}

func TestBubbleSortWorstCase(t *testing.T) {
	els := []int{9, 8, 7, 6, 5}

	els = BubbleSort(els)

	if els == nil {
		t.Error("Not expecting els to be nil")
	}

	if len(els) != 5 {
		t.Error("Expected 5 elements in return")
	}

	if els[0] != 5 {
		t.Error("Expected 5 for element")
	}

	if els[2] != 7 {
		t.Error("Expected 7 for element")
	}

	if els[4] != 9 {
		t.Error("Expected 9 for element")
	}

}

func TestBubbleSortBestCase(t *testing.T) {
	els := []int{5, 6, 7, 8, 9}

	els = BubbleSort(els)

	if els == nil {
		t.Error("Not expecting els to be nil")
	}

	if len(els) != 5 {
		t.Error("Expected 5 elements in return")
	}

	if els[0] != 5 {
		t.Error("Expected 5 for element")
	}

	if els[2] != 7 {
		t.Error("Expected 7 for element")
	}

	if els[4] != 9 {
		t.Error("Expected 9 for element")
	}

}

// functs to benchmark
func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements(t *testing.T) {
	els := getElements(5)
	if els == nil {
		t.Error("Not expecting els to be nil")
	}

	if len(els) != 5 {
		t.Error("els should have 5 elements")
	}

	if els[4] != 0 {
		t.Error("Expected els[4] to be 0")
	}
}

func BenchmarkBubbleSort10(b *testing.B) {
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	els := getElements(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}
