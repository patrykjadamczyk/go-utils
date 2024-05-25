//go:build (js && wasm)

package os

const (
	OS_GOOS = "js"
	OS_GOARCH = "wasm"
)