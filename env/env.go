package env

import (
	"github.com/gookit/goutil/envutil"
)

func GetEnv(key string) string {
    return envutil.Getenv(key)
}

func SetEnv(key string, value string) {
    envutil.SetEnvMap(map[string]string{key: value})
}

func IsGitHubActions() bool {
    return envutil.IsGithubActions()
}

func OsIsLinux() bool {
    return envutil.IsLinux()
}

func OsIsMac() bool {
    return envutil.IsMac()
}

func OsIsMSys() bool {
    return envutil.IsMSys()
}

func OsIsWsl() bool {
    return envutil.IsWSL()
}

func OsIsWindows() bool {
    return envutil.IsWindows()
}
