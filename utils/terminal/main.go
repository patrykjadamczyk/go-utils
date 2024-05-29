package terminal

import (
	"os"

	"github.com/patrykjadamczyk/go-utils/math"
	"golang.org/x/term"
)

// Get Terminal Size
func GetTerminalSize() math.Vector2[int] {
	fd := int(os.Stdout.Fd())
	width, height, err := term.GetSize(fd)
	if err != nil {
		return math.Vector2[int]{X: 80, Y: 24}
	}
	return math.Vector2[int]{X: width, Y: height}
}

// Calculate vw unit for terminal
func VW(vwsize int) int {
	tsw := GetTerminalSize().X

	if vwsize == 100 {
		return tsw
	}

	ftsw := float32(tsw)
	fvwsize := float32(vwsize)
	res := (ftsw / 100.0) * fvwsize
	return int(res)
}

// Calculate vh unit for terminal
func VH(vhsize int) int {
	tsh := GetTerminalSize().Y

	if vhsize == 100 {
		return tsh
	}

	ftsh := float32(tsh)
	fvhsize := float32(vhsize)
	res := (ftsh / 100.0) * fvhsize
	return int(res)
}
