package chefondate_test

import (
	"testing"

	chefondate "github.com/cfabrica46/codechef/chefondate/mysolution"
)

func TestChefOnDate(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		subtestingName string
		result         string
		inX            int
		inY            int
	}{
		{
			subtestingName: "OK",
			inX:            4,
			inY:            0,
			result:         "",
		},
		{
			subtestingName: "OK",
			inX:            1,
			inY:            1,
			result:         "YES",
		},
		{
			subtestingName: "OK",
			inX:            1,
			inY:            2,
			result:         "NO",
		},
		{
			subtestingName: "OK",
			inX:            2,
			inY:            1,
			result:         "YES",
		},
		{
			subtestingName: "OK",
			inX:            50,
			inY:            100,
			result:         "NO",
		},
	} {
		tt := tt
		t.Run(tt.subtestingName, func(t *testing.T) {
			t.Parallel()

			result := chefondate.ChefOnDate(tt.inX, tt.inY)

			if result != tt.result {
				t.Errorf("expected: %s; obtained: %s", tt.result, result)
			}
		})
	}
}
