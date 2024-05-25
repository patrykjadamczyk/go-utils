package os

import "slices"

const (
	OS_IS_AIX = OS_GOOS == "aix"
	OS_IS_ANDROID = OS_GOOS == "android"
	OS_IS_DARWIN = OS_GOOS == "darwin"
	OS_IS_DRAGONFLY = OS_GOOS == "dragonfly"
	OS_IS_FREEBSD = OS_GOOS == "freebsd"
	OS_IS_ILLUMOS = OS_GOOS == "illumos"
	OS_IS_IOS = OS_GOOS == "ios"
	OS_IS_JS = OS_GOOS == "js"
	OS_IS_LINUX = OS_GOOS == "linux"
	OS_IS_NETBSD = OS_GOOS == "netbsd"
	OS_IS_OPENBSD = OS_GOOS == "openbsd"
	OS_IS_PLAN9 = OS_GOOS == "plan9"
	OS_IS_SOLARIS = OS_GOOS == "solaris"
	OS_IS_WASIP1 = OS_GOOS == "wasip1"
	OS_IS_WINDOWS = OS_GOOS == "windows"
)

const (
	OSARCH_IS_386 = OS_GOARCH == "386"
	OSARCH_IS_PPC64 = OS_GOARCH == "ppc64"
	OSARCH_IS_AMD64 = OS_GOARCH == "amd64"
	OSARCH_IS_ARM = OS_GOARCH == "arm"
	OSARCH_IS_ARM64 = OS_GOARCH == "arm64"
	OSARCH_IS_RISCV64 = OS_GOARCH == "riscv64"
	OSARCH_IS_WASM = OS_GOARCH == "wasm"
	OSARCH_IS_LOONG64 = OS_GOARCH == "loong64"
	OSARCH_IS_MIPS = OS_GOARCH == "mips"
	OSARCH_IS_MIPS64 = OS_GOARCH == "mips64"
	OSARCH_IS_MIPS64LE = OS_GOARCH == "mips64le"
	OSARCH_IS_MIPSLE = OS_GOARCH == "mipsle"
	OSARCH_IS_S390X = OS_GOARCH == "s390x"
)

func RunOnlyOnOS(f func(), oses ...string) {
	if slices.Contains(oses, OS_GOOS) {
		f()
	}
}

func RunOnlyOnARCH(f func(), arches ...string) {
	if slices.Contains(arches, OS_GOARCH) {
		f()
	}
}
