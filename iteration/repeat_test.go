package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("run 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		AssertRepeat(t, repeated, expected)
	})

	t.Run("run 10 times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		AssertRepeat(t, repeated, expected)

	})

}

func AssertRepeat(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Fatalf("want: %q, got: %q", repeated, expected)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
