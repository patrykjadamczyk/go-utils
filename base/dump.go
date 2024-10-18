package base

import (
	"github.com/gookit/goutil/dump"
)

// defaultSkip is the default number of stack frames to skip
// This constant is meant to be directly copied from the source code line below
// ~/go/pkg/mod/github.com/gookit/goutil@<ver>/dump/dump.go:24
const defaultSkip = 3

// This constant is meant to be tested manually and changed so it shows correct information
const DevDumpSkip = defaultSkip + 0

// Dump variable to console output
func Dump(vs ...any) {
	dump.NewWithOptions(
		dump.WithCallerSkip(DevDumpSkip),
	).Print(vs...)
}

// Dump variable to console output with specified depth
func DumpDepth(depth int, vs ...any) {
	dump.NewWithOptions(
		func(opts *dump.Options) {
			opts.CallerSkip = DevDumpSkip
			opts.MaxDepth = depth
		},
	).Print(vs...)
}

// Dump variable to console output with unlimited depth
func DumpAll(vs ...any) {
	dump.NewWithOptions(
		func(opts *dump.Options) {
			opts.CallerSkip = DevDumpSkip
			opts.MaxDepth = GetMaxForType[int]()
		},
	).Print(vs...)
}
