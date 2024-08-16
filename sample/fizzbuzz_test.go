package sample

import (
	"fmt"
	"testing"
)

func TestFizzbuzz(t *testing.T) {
	testCases := []struct {
		n    int
		want string
	}{
		{n: 1, want: "1"},
		{n: 2, want: "2"},
		{n: 3, want: "Fizz"},
		{n: 4, want: "4"},
		{n: 5, want: "Buzz"},
		{n: 6, want: "Fizz"},
		{n: 7, want: "7"},
		{n: 8, want: "8"},
		{n: 9, want: "Fizz"},
		{n: 10, want: "Buzz"},
		{n: 11, want: "11"},
		{n: 12, want: "Fizz"},
		{n: 13, want: "13"},
		{n: 14, want: "14"},
		{n: 15, want: "FizzBuzz"},
	}

	for _, tC := range testCases {
		tC := tC
		desc := fmt.Sprintf(`Convert(%v)`, tC.n)

		t.Run(desc, func(t *testing.T) {
			t.Parallel()
			testConvert(t, tC.n, tC.want)
		})
	}
}

func testConvert(t *testing.T, n int, want string) {
	t.Helper()
	got := Convert(n)

	if got != want {
		t.Errorf(`Convert(%v) = %q but want %q`, n, got, want)
	}
}
