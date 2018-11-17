package session

import (
	"errors"
)

var (
	ErrSessionNotExist = errors.New("session not exist")
	ErrKeyNotExist     = errors.New("key not exist in session")
)
