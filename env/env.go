package env

import (
	"github.com/gookit/goutil/envutil"
)

// Get Environment Variable for specified key
func GetEnv(key string) string {
    return envutil.Getenv(key)
}

// Set Environment Variable for specified key with specified value
func SetEnv(key string, value string) {
    envutil.SetEnvMap(map[string]string{key: value})
}

// Check if current environment is GitHub Actions
func IsGitHubActions() bool {
    return envutil.IsGithubActions()
}

// Check if current environment is Linux
func OsIsLinux() bool {
    return envutil.IsLinux()
}

// Check if current environment is Mac
func OsIsMac() bool {
    return envutil.IsMac()
}

// Check if current environment is MSys
func OsIsMSys() bool {
    return envutil.IsMSys()
}

// Check if current environment is WSL
func OsIsWsl() bool {
    return envutil.IsWSL()
}

// Check if current environment is Windows
func OsIsWindows() bool {
    return envutil.IsWindows()
}
