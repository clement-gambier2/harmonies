package main

// Token represents a colored token on the board
type Token struct {
	Color  TokenColor
	Height TokenHeight
	Cube   bool // Whether an animal cube is on this token
}

// Token colors
type TokenColor int

const (
	Empty TokenColor = iota
	Gray
	Blue
	Brown
	Green
	Yellow
	Red
)

// Token heights
type TokenHeight int

const (
	NoHeight TokenHeight = iota
	OneHigh
	TwoHigh
	ThreeHigh
)

// Define the color names for display
func ColorName(color TokenColor) string {
	switch color {
	case Empty:
		return "Empty"
	case Gray:
		return "Gray"
	case Blue:
		return "Blue"
	case Brown:
		return "Brown"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	case Red:
		return "Red"
	default:
		return "Unknown"
	}
}

// Define ANSI color codes for terminal display
func ColorCode(color TokenColor) string {
	switch color {
	case Empty:
		return "\033[0m" // Reset
	case Gray:
		return "\033[30;47m" // Black on white
	case Blue:
		return "\033[37;44m" // White on blue
	case Brown:
		return "\033[30;43m" // Black on yellow (approximation for brown)
	case Green:
		return "\033[30;42m" // Black on green
	case Yellow:
		return "\033[30;103m" // Black on bright yellow
	case Red:
		return "\033[37;41m" // White on red
	default:
		return "\033[0m" // Reset
	}
}
