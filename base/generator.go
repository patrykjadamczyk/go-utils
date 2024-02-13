package base

// Make Generator
// abortChannel is used to stop generator
// generatorFunc is generator function that is called every iteration like next. Have to have parameter for channel for
// writing values and returning bool if it should stop (false) or continue (true)
// Note that this generator is executed in goroutine
func MakeGenerator[T any](
	abortChannel <-chan struct{},
	generatorFunc func(chan T) bool,
) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for {
			select {
			case <-abortChannel:
				return
			default:
				ok := generatorFunc(ch)
				if !ok {
					return
				}
			}
		}
	}()
	return ch
}

// Yield in Generator
// Add Value to Channel
func Yield[T any](ch chan<- T, value T) {
	ch <- value
}
