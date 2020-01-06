package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCase(t *testing.T) {
	// Initialization:
	els := []int{9, 8, 7, 6, 5}

	// Execution:
	BubbleSort(els)

	// Validation:
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4])
}

func TestBubbleSortBestCase(t *testing.T) {
	// Initialization:
	els := []int{5, 6, 7, 8, 9}

	// Execution:
	BubbleSort(els)

	// Validation:
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4])
}

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

	assert.NotNil(t, els)

	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 4, els[1])
	assert.EqualValues(t, 3, els[2])
	assert.EqualValues(t, 2, els[3])
	assert.EqualValues(t, 1, els[4])
}

func BenchmarkBubbleSort10(b *testing.B) {
	els := getElements(10)
	for i := 10; i < b.N; i++ {
		BubbleSort(els)
	}
}
