package errors

import "errors"

var (
	DatabaseInternalError  = errors.New("internal error")
	DatabaseRecordNotFound = errors.New("record not found")
)

var (
	Unauthorized = errors.New("you do not have permission to access this resource")
	Forbidden    = errors.New("you do not have permission to access this resource")
)

var (
	UsernameOrPasswordNotMatch = errors.New("username or password not match")
)

var (
	PostNotFound = errors.New("post not found")
)

var (
	UserNotFound = errors.New("user not found")
)
