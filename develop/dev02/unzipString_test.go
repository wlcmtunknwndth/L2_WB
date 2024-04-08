package unzipString

import (
	"fmt"
	"testing"
)

func TestUnzipString(t *testing.T) {
	testCases := []struct {
		requested string
		ans       string
	}{
		{
			requested: "a4bc2d5e",
			ans:       "aaaabccddddde",
		},
		{
			requested: "abcd",
			ans:       "abcd",
		},
		{
			requested: "45",
			ans:       "",
		},
		{
			requested: "",
			ans:       "",
		},
		{
			requested: "tsfds45gf",
			ans:       "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.requested, func(t *testing.T) {
			unzipped := UnzipString(tc.requested)
			if tc.ans != unzipped {
				t.Errorf("uneven: %s and %s: ", tc.ans, unzipped)
			}
			fmt.Printf("Exp: %s\n", tc.ans)
			fmt.Printf("Got: %s\n", unzipped)
		})
	}
}
