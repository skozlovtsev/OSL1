package console

const (
	RedText    = "\u001b[31m"
	GreenText  = "\u001b[32m"
	YellowText = "\u001b[33m"
	BlueText   = "\u001b[34m"

	RedBackground    = "\u001b[41m;1m"
	GreenBackground  = "\u001b[42m;1m"
	YellowBackground = "\u001b[43m;1m"
	BlueBackground   = "\u001b[44m;1m"

	// Decorators
	Bold      = "\u001b[1m"
	Underline = "\u001b[4m"
	Reversed  = "\u001b[7m" // Reversed selection

	ResetColor  = "\u001b[0m"
	ClearLine   = "\u001b[2K" // Clear entire line
	ClearScreen = "\u001b[2J" // Clear entire screen
)
