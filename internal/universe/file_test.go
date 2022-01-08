package universe_test

import (
	"conway/internal/universe"
	"fmt"
	"testing"
)

func TestUniverse_FromFile(t *testing.T) {
	tests := []struct {
		path string
		want universe.Universe
	}{
		{
			"fixtures/basic.cells",
			universe.Universe{
				{true, false},
				{false, true},
			},
		},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("From fixture %s", tt.path)
		t.Run(testname, func(t *testing.T) {
			got, err := universe.FromFile(tt.path)
			if err != nil {
				panic(err)
			}

			if !got.Equals(tt.want) {
				t.Errorf("\ngot: \n%v\nwant: \n%v\n", got, tt.want)
			}
		})
	}
}
