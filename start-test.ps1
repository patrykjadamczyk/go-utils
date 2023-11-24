#!/usr/bin/env pwsh
Clear-Host
go install gotest.tools/gotestsum@latest
gotestsum --watch  -- -v -count=1 ./...
