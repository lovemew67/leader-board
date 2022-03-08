package lb

import (
	"fmt"
	"strconv"
)

const (
	DefaultMaxLengthInt = 10
)

var (
	DefaultMaxLengthStr = strconv.FormatInt(DefaultMaxLengthInt, 10)
)

const (
	ErrRedisKeyNotFoundStr = "redigo: nil returned"
)

var (
	ErrRedisKeyNotFound = fmt.Errorf(ErrRedisKeyNotFoundStr)
)
