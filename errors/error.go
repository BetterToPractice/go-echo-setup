package errors

import "errors"

var (
	DatabaseInternalError  = errors.New("database internal error")
	DatabaseRecordNotFound = errors.New("database record not found")
	RecordNotFound         = errors.New("record not found")
)

var (
	Unauthorized = errors.New("unauthorized: You do not have permission to access this resource")
	Forbidden    = errors.New("forbidden: You do not have permission to access this resource")
)

var (
	UsernameOrPasswordNotMatch = errors.New("username or password not match")
)
