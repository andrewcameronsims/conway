package command

type Command int

const (
	None Command = iota
	Forward
	Back
	TogglePlay
	Exit
)
