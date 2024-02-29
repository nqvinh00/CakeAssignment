package model

type ErrorCode int

const (
	Success ErrorCode = iota
	UnknownError
)

var ErrorMessage = map[ErrorCode]string{}
