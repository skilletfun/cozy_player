package utils

import (
	"testing"
)

func TestSumSliceInts(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		sumfunc func(int) int
		result  int
	}{
		{"slice with zeros", []int{0, 0, 0, 0, 0}, func(e int) int { return e }, 0},
		{"slice 0, 1, 2, 3, 4", []int{0, 1, 2, 3, 4}, func(e int) int { return e }, 10},
		{"slice 0, -1, -2, -3, -4", []int{0, -1, -2, -3, -4}, func(e int) int { return e }, -10},
		{"slice 0, -1, -2, -3, -4", []int{0, 1, -2, 3, -4}, func(e int) int { return e }, -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumSlice(tt.slice, tt.sumfunc); got != tt.result {
				t.Errorf("SumSlice(%v, ...) = %v, want %v", tt.slice, got, tt.result)
			}
		})
	}
}

type TestStruct interface {
	Get() int
}

type testStruct struct {
	data int
}

func (s testStruct) Get() int {
	return s.data
}

func TestSumSliceStrings(t *testing.T) {
	tests := []struct {
		name    string
		slice   []TestStruct
		sumfunc func(e TestStruct) int
		result  int
	}{
		{"slice with zeros", []TestStruct{testStruct{0}, testStruct{0}, testStruct{0}}, func(e TestStruct) int { return e.Get() }, 0},
		{"slice with zeros", []TestStruct{testStruct{1}, testStruct{1}, testStruct{1}}, func(e TestStruct) int { return e.Get() }, 3},
		{"slice with zeros", []TestStruct{testStruct{-1}, testStruct{-1}, testStruct{-1}}, func(e TestStruct) int { return e.Get() }, -3},
		{"slice with zeros", []TestStruct{testStruct{10}, testStruct{-5}, testStruct{-5}}, func(e TestStruct) int { return e.Get() }, 0},
		{"slice with zeros", []TestStruct{testStruct{100}, testStruct{100}, testStruct{100}}, func(e TestStruct) int { return e.Get() }, 300},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumSlice(tt.slice, tt.sumfunc); got != tt.result {
				t.Errorf("SumSlice(%v, ...) = %v, want %v", tt.slice, got, tt.result)
			}
		})
	}
}