package universe

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// FromFile will build a Universe from any file that
// follows the .cells plaintext format specified at
// https://www.conwaylife.com/wiki/Plaintext
func FromFile(path string) (Universe, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	pattern, err := parse(file)
	if err != nil {
		return nil, fmt.Errorf("could not parse file: %v", err)
	}

	return pattern, nil
}

func parse(file *os.File) ([][]bool, error) {
	var pattern [][]bool

	scanner := bufio.NewScanner(file)
	// validate equal row length too

	var i int
	var rowLen int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")

		// Ignore comments
		if line[0] == "!" {
			continue
		}

		if rowLen != 0 && len(line) != rowLen {
			return nil, errors.New("all rows must be of equal length")
		}

		var row []bool
		for j := 0; j < len(line); j++ {
			switch line[j] {
			case "O":
				row = append(row, true)
			case ".":
				row = append(row, false)
			default:
				return nil, fmt.Errorf("invalid character %v", line[j])
			}
		}

		pattern = append(pattern, row)
		rowLen = len(row)
		i++
	}

	return pattern, nil
}
