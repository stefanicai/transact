package main

// generate client.go and server API
//go:generate go run github.com/ogen-go/ogen/cmd/ogen --target ./internal/api --package api --clean ./swagger/transact.yaml
