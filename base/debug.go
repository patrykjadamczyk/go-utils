package base

var DebugMode bool = false

func SetDebugMode(mode bool) {
	DebugMode = mode
}

func GetDebugMode() bool {
	return DebugMode
}

func ExecuteInDebugMode(f func()) {
	if DebugMode {
		f()
	}
}

func ExecuteAndPassthrough(f func([]any) []any, values ...any) []any {
	return f(values)
}
