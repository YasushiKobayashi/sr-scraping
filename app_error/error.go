package app_error

import (
	"fmt"
	"runtime"
)

type (
	Causer interface {
		Cause() error
	}

	AppError struct {
		Err            error
		CliCode        int
		ProgramCounter uintptr
	}
)

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	} else {
		return "AppError"
	}
}

func FindAppError(err error) *AppError {
	_err := err
	for _err != nil {
		if ue, ok := _err.(*AppError); ok {
			return ue
		}

		cause, ok := _err.(Causer)
		if !ok {
			break
		}
		_err = cause.Cause()
	}

	return nil
}

func newAppError(err error, code int) *AppError {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("err")
	}
	return &AppError{
		Err:            err,
		CliCode:        code,
		ProgramCounter: pc,
	}
}

func NewBadRequestErr(err error) *AppError {
	return newAppError(err, 128)
}

func NewErr(err error) *AppError {
	return newAppError(err, 255)
}
