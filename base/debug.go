package base

// DebugMode Variable
var DebugMode bool = false

// Set Debug Mode
func SetDebugMode(mode bool) {
	DebugMode = mode
}

// Get Current State of Debug Mode
func GetDebugMode() bool {
	return DebugMode
}

// Execute specified function only in Debug Mode
func ExecuteInDebugMode(f func()) {
	if DebugMode {
		f()
	}
}

// Execute specified function and pass through all values
func ExecuteAndPassthrough(f func([]any) []any, values ...any) []any {
	return f(values)
}
