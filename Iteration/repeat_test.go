package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	tests := []struct {
		name      string
		character string
		number    int
		expected  string
	}{
		{"repeat once", "a", 1, "a"},
		{"repeat three times", "b", 3, "bbb"},
		{"repeat zero times", "c", 0, ""},
		{"repeat multi-char string", "ab", 2, "abab"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Repeat(tt.character, tt.number)
			if result != tt.expected {
				t.Errorf(
					"Repeat(%q, %d) = %q; want %q",
					tt.character,
					tt.number,
					result,
					tt.expected,
				)
			}
		})
	}
}

 

func BenchmarkRepeat(b *testing.B) {
	i := 1
	for b.Loop() {
		Repeat("a", i)
	}
}

func ExampleRepeat() {
	i := 5
	repeated := Repeat("a", i)
	fmt.Println(repeated)
	//Output: aaaaa
}
