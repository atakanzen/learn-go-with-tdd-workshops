package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	want := "aaaaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestRepeatWithSb(t *testing.T) {
	got := RepeatWithSb("a", 6)
	want := "aaaaaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func BenchmarkRepeatWithSb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatWithSb("a", 6)
	}
}

func ExampleRepeat() {
	repeated := Repeat("ab")
	fmt.Println(repeated)
	// Output: ababababab
}

func ExampleRepeatWithSb() {
	repeated := RepeatWithSb("ab", 2)
	fmt.Println(repeated)
	// Output: abab
}
