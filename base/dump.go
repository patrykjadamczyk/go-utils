package base

import (
	"github.com/gookit/goutil/dump"
)

// Dump variable to console output
func Dump(vs ...any) {
	dump.P(vs...)
}

// Dump variable to console output with specified depth
func DumpDepth(depth int, vs ...any) {
	dump.NewWithOptions(
		func(opts *dump.Options) {
			opts.MaxDepth = depth
		},
	).Print(vs...)
}

// Dump variable to console output with unlimited depth
func DumpAll(vs ...any) {
	DumpDepth(GetMaxForType[int](), vs...)
}
