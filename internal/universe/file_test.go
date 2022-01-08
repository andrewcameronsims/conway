package universe_test

import (
	"conway/internal/universe"
	"testing"
)

func TestUniverse_FromFile(t *testing.T) {
	tests := []struct {
		desc   string
		path   string
		uni    universe.Universe
		errMsg string
	}{
		{
			"It creates a universe from file",
			"fixtures/basic.cells",
			universe.Universe{
				{true, false},
				{false, true},
			},
			"",
		},
		{
			"It ignores the header of the file",
			"fixtures/header.cells",
			universe.Universe{
				{true, false},
				{false, true},
			},
			"",
		},
		{
			"It returns an error when the data contains invalid dimensions",
			"fixtures/invalid_dimensions.cells",
			nil,
			"could not parse file: all rows must be of equal length",
		},
		{
			"It returns an error when the data contains invalid characters",
			"fixtures/invalid_characters.cells",
			nil,
			"could not parse file: invalid character #",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got, err := universe.FromFile(tt.path)

			if err != nil && err.Error() != tt.errMsg {
				t.Errorf("\ngot: \n%v\nwant: \n%v\n", err, tt.errMsg)
			}

			if !got.Equals(tt.uni) {
				t.Errorf("\ngot: \n%v\nwant: \n%v\n", got, tt.uni)
			}
		})
	}
}
