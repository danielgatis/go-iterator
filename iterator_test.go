package iterator_test

import (
	"reflect"
	"testing"

	"github.com/danielgatis/go-iterator"
)

func TestIteratorImpl_GetNext(t *testing.T) {
	tests := []struct {
		name     string
		items    []int
		expected []int
	}{
		{
			name:     "Empty slice",
			items:    []int{},
			expected: []int{},
		},
		{
			name:     "Single item slice",
			items:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Multiple item slice",
			items:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iterator := iterator.NewIterator(tt.items)
			result := make([]int, 0, len(tt.items))

			for i := 0; i < len(tt.items); i++ {
				item, ok := iterator.GetNext()
				if !ok {
					t.Errorf("getNext() returned false, expected true")
				}
				result = append(result, item)
			}

			if _, ok := iterator.GetNext(); ok {
				t.Errorf("getNext() returned true after all items were retrieved, expected false")
			}

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestIteratorImpl_GetNextOrDefault(t *testing.T) {
	tests := []struct {
		name         string
		items        []int
		defaultValue int
		expected     []int
	}{
		{
			name:         "Empty slice",
			items:        []int{},
			defaultValue: 0,
			expected:     []int{0},
		},
		{
			name:         "Single item slice",
			items:        []int{1},
			defaultValue: 0,
			expected:     []int{1, 0},
		},
		{
			name:         "Multiple item slice",
			items:        []int{1, 2, 3, 4, 5},
			defaultValue: 0,
			expected:     []int{1, 2, 3, 4, 5, 0},
		},
		{
			name:         "Empty slice with default value",
			items:        []int{},
			defaultValue: 1,
			expected:     []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iterator := iterator.NewIterator(tt.items)
			result := make([]int, 0, len(tt.items))

			for i := 0; i < len(tt.items)+1; i++ {
				item := iterator.GetNextOrDefault(tt.defaultValue)
				result = append(result, item)
			}

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
