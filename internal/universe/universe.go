package universe

type Universe [][]bool

type Location struct {
	row int
	col int
}

func New(size int) Universe {
	var universe Universe

	for i := 0; i < size; i++ {
		var row []bool
		for j := 0; j < size; j++ {
			row = append(row, false)
		}
		universe = append(universe, row)
	}

	return universe
}

func (u Universe) Tick() Universe {
	var next [][]bool

	for i, row := range u {
		var nextRow []bool

		for j, cell := range row {
			loc := Location{row: i, col: j}
			live := u.liveNeighbours(loc)

			if lifeConditions(cell, live) {
				nextRow = append(nextRow, true)
			} else {
				nextRow = append(nextRow, false)
			}
		}

		next = append(next, nextRow)
	}

	return next
}

func (u Universe) Equals(other Universe) bool {
	if len(u) != len(other) {
		return false
	}

	for i, row := range u {
		for j, cell := range row {
			if cell != other[i][j] {
				return false
			}
		}
	}

	return true
}

func (u Universe) String() string {
	uStr := ""

	for _, row := range u {
		rowStr := ""
		for _, cell := range row {
			if cell {
				rowStr += "O"
			} else {
				rowStr += "."
			}
		}
		rowStr += "\n"
		uStr += rowStr
	}

	return uStr
}

func lifeConditions(cell bool, live int) bool {
	return !cell && live == 3 || (cell && (live == 2 || live == 3))
}

func (u Universe) get(loc Location) bool {
	return u[loc.row][loc.col]
}

func (u Universe) contains(loc Location) bool {
	if loc.row < 0 || loc.row > len(u)-1 {
		return false
	}

	if loc.col < 0 || loc.col > len(u[0])-1 {
		return false
	}

	return true
}

func (u Universe) liveNeighbours(loc Location) int {
	var live int

	for _, n := range u.neighbours(loc) {
		if u.get(n) {
			live++
		}
	}

	return live
}

func (u Universe) neighbours(loc Location) []Location {
	var ns []Location
	adjacent := []Location{
		{row: loc.row - 1, col: loc.col - 1},
		{row: loc.row - 1, col: loc.col},
		{row: loc.row - 1, col: loc.col + 1},
		{row: loc.row, col: loc.col - 1},
		{row: loc.row, col: loc.col + 1},
		{row: loc.row + 1, col: loc.col - 1},
		{row: loc.row + 1, col: loc.col},
		{row: loc.row + 1, col: loc.col + 1},
	}

	for i := 0; i < len(adjacent); i++ {
		if u.contains(adjacent[i]) {
			ns = append(ns, adjacent[i])
		}
	}

	return ns
}
