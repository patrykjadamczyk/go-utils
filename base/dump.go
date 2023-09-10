package base

import (
	"github.com/gookit/goutil/dump"
)

// Dump variable to console output
func Dump(vs ...any) {
	dump.P(vs...)
}
