#!/usr/bin/env nu
let GOARCH = (^go tool dist list) | split row "\n"
$GOARCH | each { |goarch|
    let go_os = (($goarch | split row "/") | first)
    let go_arch = (($goarch | split row "/") | last)
    []
        | append [ $"//go:build \(($go_os) && ($go_arch)\)" ""]
        | append [ "package os" ""]
        | append [ "const (" ]
        | append [ $"	OS_GOOS = \"($go_os)\"" ]
        | append [ $"	OS_GOARCH = \"($go_arch)\"" ]
        | append [ ")"]
        | str join "\n"
        | save -f $"./main_($go_os)_($go_arch).go"
}
