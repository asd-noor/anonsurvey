package consts

import (
	"anonsurvey/types"
)

const (
	S_READ types.Permission = 1 << iota
	S_WRITE
)

const (
	S_RW = S_READ | S_WRITE
)
