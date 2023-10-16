package errors

import "errors"

var (
	DatabaseInternalError  = errors.New("database internal error")
	DatabaseRecordNotFound = errors.New("database record not found")
	RecordNotFound         = errors.New("record not found")
)

var (
	UsernameOrPasswordNotMatch = errors.New("username or password not match")
)
