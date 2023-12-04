package mocks

import "github.com/stefanicai/transact/internal/api"

func OptString(s string) api.OptString {
	return api.OptString{
		Value: s,
		Set:   true,
	}
}
