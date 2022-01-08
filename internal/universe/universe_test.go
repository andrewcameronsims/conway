package universe_test

import (
	"conway/internal/universe"
	"fmt"
	"testing"
)

func TestUniverse_New(t *testing.T) {
	tests := []struct {
		size int
		want universe.Universe
	}{
		{
			2,
			universe.Universe{
				{false, false},
				{false, false},
			},
		},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("Universe of size %d", tt.size)
		t.Run(testname, func(t *testing.T) {
			got := universe.New(tt.size)
			if !got.Equals(tt.want) {
				t.Errorf("\ngot: \n%v\nwant: \n%v\n", got, tt.want)
			}
		})
	}
}

func TestUniverse_Tick(t *testing.T) {
	var tests = []struct {
		name string
		init universe.Universe
		want universe.Universe
	}{
		{
			"Blinker",
			universe.Universe{
				{false, true, false},
				{false, true, false},
				{false, true, false},
			},
			universe.Universe{
				{false, false, false},
				{true, true, true},
				{false, false, false},
			},
		},
		{
			"Block",
			universe.Universe{
				{true, true},
				{true, true},
			},
			universe.Universe{
				{true, true},
				{true, true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.init.Tick()

			if !got.Equals(tt.want) {
				t.Errorf("\ngot: \n%v\nwant: \n%v\n", got, tt.want)
			}
		})
	}
}
