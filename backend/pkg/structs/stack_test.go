package structs

import "testing"

func TestStackSize(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"stack 0 elements", 0},
		{"stack 1 element", 1},
		{"stack 2 elements", 2},
		{"stack 3 elements", 3},
		{"stack 4 elements", 4},
		{"stack 5 elements", 5},
		{"stack 10 elements", 10},
		{"stack 25 elements", 25},
		{"stack 100 elements", 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack[any]()

			for i := 0; i < tt.size; i++ {
				stack.Push(i)
			}

			if got := stack.Size(); got != tt.size {
				t.Errorf("stack.Size() = %v, want %v", got, tt.size)
			}
		})
	}
}

func TestStackPopFromEmptyStack(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Code didn't panic")
		}
	}()

	stack := NewStack[any]()
	stack.Pop()
}

func TestStackIsEmpty(t *testing.T) {
	t.Run("check default stack", func(t *testing.T) {
		stack := NewStack[any]()
		if empty := stack.IsEmpty(); !empty {
			t.Errorf("stack.IsEmpty() = %v, want %v", empty, true)
		}
	})

	t.Run("check stack not empty", func(t *testing.T) {
		stack := NewStack[any]()
		stack.Push(1)
		if empty := stack.IsEmpty(); empty {
			t.Errorf("stack.IsEmpty() = %v, want %v", empty, false)
		}
	})

	t.Run("check stack empty", func(t *testing.T) {
		stack := NewStack[any]()
		stack.Push(1)
		stack.Pop()
		if empty := stack.IsEmpty(); !empty {
			t.Errorf("stack.IsEmpty() = %v, want %v", empty, true)
		}
	})
}

func TestStackPushPopRepeat(t *testing.T) {
	stack := NewStack[int]()

	for i := range 100 {
		stack.Push(i)

		if got := stack.Pop(); got != i {
			t.Errorf("stack.Pop() = %v, want %v", got, i)
		}
	}
}

func TestStackPushPop(t *testing.T) {
	stack := NewStack[int]()

	for i := range 100 {
		stack.Push(i)
	}

	for i := 99; i >= 0; i-- {
		if got := stack.Pop(); got != i {
			t.Errorf("stack.Pop() = %v, want %v", got, i)
		}
	}
}
