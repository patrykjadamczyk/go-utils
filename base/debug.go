package base

// DebugMode Variable
var DebugMode Observable[bool] = MakeObservable[bool](false)

// Set Debug Mode
func SetDebugMode(mode bool) {
	DebugMode.SetValue(mode)
}

// Get Current State of Debug Mode
func GetDebugMode() bool {
	return DebugMode.GetValue()
}

// Execute specified function only in Debug Mode
func ExecuteInDebugMode(f func()) {
	if GetDebugMode() {
		f()
	}
}

// Execute specified function and pass through all values
func ExecuteAndPassthrough(f func([]any) []any, values ...any) []any {
	return f(values)
}
