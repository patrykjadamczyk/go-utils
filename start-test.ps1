#!/usr/bin/env pwsh
Clear-Host
go install gotest.tools/gotestsum@latest
gotestsum -- -v -count=1 ./...
gotestsum --watch  -- -v -count=1 ./...
