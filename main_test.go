package main

import "testing"

func TestIndex(t *testing.T) {
	tests := []struct {
		name      string
		elements  []int
		predicate func(int) bool
		want      int
	}{
		{"Index", []int{1, 2, 3, 4, 5}, func(i int) bool {
			return i == 5
		}, 4},
		{"No Index", []int{1, 2, 3, 4, 5}, func(i int) bool {
			return i == 6
		}, -1},
	}
	for _, test := range tests {
		if got := Index(test.elements, test.predicate); got != test.want {
			t.Error("test for", test.name, "got:", got, "want:", test.want)
		}
	}
}
func TestAll(t *testing.T) {
	tests := []struct {
		name      string
		elements  []int
		predicate func(int) bool
		want      bool
	}{
		{"All true", []int{1, 2, 3, 4, 5}, func(i int) bool {
			return i == 5
		},

			false},
		{"All false", []int{3, 3, 3, 3, 3}, func(i int) bool {
			return i == 3
		},
			true},}
	for _, test := range tests {
		if got := All(test.elements, test.predicate); got != test.want {
			t.Error("test for", test.name, "got:", got, "want:", test.want)
		}
	}
}
func TestAny(t *testing.T) {
	tests := []struct {
		name      string
		elements  []int
		predicate func(int) bool
		want      bool
	}{
		{"Any true", []int{1, 2, 3, 4, 5}, func(i int) bool {
			return i == 1
		}, true},
		{"Any false", []int{1, 2, 3, 4, 5}, func(i int) bool {
			return i == 7
		}, false},
	}
	for _, test := range tests {
		if got := Any(test.elements, test.predicate); got != test.want {
			t.Error("test for", test.name, "got:", got, "want:", test.want)
		}
	}
}
func TestNone(t *testing.T) {
	tests := []struct {
		name      string
		elements  []int
		predicate func(int) bool
		want      bool
	}{
		{"None true", []int{1, 2, 3, 4, 5}, func(i int) bool {
			return i == 6
		}, true},
		{"None false", []int{1, 2, 3, 4, 5}, func(i int) bool {
			return i == 2
		}, false},
	}
	for _, test := range tests {
		if got := None(test.elements, test.predicate); got != test.want {
			t.Error("test for", test.name, "got:", got, "want:", test.want)
		}
	}
}
func TestFind(t *testing.T) {
	if Find([]int{1, 2, 3, 4, 5}, func(element int) bool {
		return element == 3
	}) != 3 {
		t.Error("")
	}
	func() {
		defer func() {
			err := recover()
			if err == nil {
				t.Error("Panic!")
			}
		}()
		Find([]int{1, 2, 3, 4, 5}, func(element int) bool {
			return element == 8
		})
	}()

	func() {
		defer func() {
			err := recover()
			if err == nil {
				t.Error("Panic!")
			}
		}()
		Find([]int{1, 2, 3, 4, 5}, func(element int) bool {
			return element == -1
		})
	}()
}
