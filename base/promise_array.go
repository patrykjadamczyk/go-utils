package base

// Run all promises and wait for their results
func PromiseAll(functions ...any) ([]any, []error) {
    promises := make([]*Promise, len(functions))
    results := make([]any, len(functions))
    errors := make([]error, len(functions))
    for i, f := range functions {
        promises[i] = MakePromisePointer(f)
        promises[i].Call()
        results[i] = promises[i].Result
        errors[i] = promises[i].Error
    }
    for _, p := range promises {
        p.Wait()
    }
    return results, errors
}
