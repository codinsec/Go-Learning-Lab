package main

import "testing"

func TestArray(t *testing.T) {
	var arr [5]int
	if len(arr) != 5 {
		t.Errorf("Expected array length 5, got %d", len(arr))
	}
	
	arr2 := [5]int{1, 2, 3, 4, 5}
	if arr2[0] != 1 || arr2[4] != 5 {
		t.Errorf("Array values incorrect")
	}
}

func TestSlice(t *testing.T) {
	slice := []int{1, 2, 3}
	if len(slice) != 3 {
		t.Errorf("Expected slice length 3, got %d", len(slice))
	}
}

func TestSliceMake(t *testing.T) {
	slice := make([]int, 3, 10)
	if len(slice) != 3 {
		t.Errorf("Expected length 3, got %d", len(slice))
	}
	if cap(slice) != 10 {
		t.Errorf("Expected capacity 10, got %d", cap(slice))
	}
}

func TestAppend(t *testing.T) {
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5)
	if len(slice) != 5 {
		t.Errorf("Expected length 5, got %d", len(slice))
	}
	if slice[3] != 4 || slice[4] != 5 {
		t.Errorf("Appended values incorrect")
	}
}

func TestCopy(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	dest := make([]int, len(source))
	copied := copy(dest, source)
	
	if copied != len(source) {
		t.Errorf("Expected %d elements copied, got %d", len(source), copied)
	}
	
	for i := range source {
		if dest[i] != source[i] {
			t.Errorf("Copy failed at index %d", i)
		}
	}
	
	// Modify source to verify independence
	source[0] = 99
	if dest[0] == 99 {
		t.Error("Dest should be independent of source")
	}
}

func TestSliceSlicing(t *testing.T) {
	original := []int{0, 1, 2, 3, 4, 5}
	slice := original[2:5]
	
	if len(slice) != 3 {
		t.Errorf("Expected length 3, got %d", len(slice))
	}
	if slice[0] != 2 || slice[2] != 4 {
		t.Errorf("Slice values incorrect")
	}
}

func TestSliceCapacity(t *testing.T) {
	slice := make([]int, 0, 5)
	if cap(slice) != 5 {
		t.Errorf("Expected capacity 5, got %d", cap(slice))
	}
	
	slice = append(slice, 1, 2, 3, 4, 5, 6) // Exceeds capacity
	if len(slice) != 6 {
		t.Errorf("Expected length 6, got %d", len(slice))
	}
	// Capacity should grow (typically doubles)
	if cap(slice) < 6 {
		t.Errorf("Capacity should grow, got %d", cap(slice))
	}
}

