package errors

import (
	stdErrors "errors"
	"fmt"
	"golang.org/x/xerrors"
)

type AppError struct {
	next    error
	message string
	frame   xerrors.Frame
}

func (e *AppError) Error() string {
	rootErr := GetRootError(e)
	if rootAppErr := AsAppError(rootErr); rootAppErr != nil {
		return rootAppErr.message
	} else {
		return rootErr.Error()
	}
}

func GetRootError(err error) error {
	appErr := AsAppError(err)
	if appErr != nil {
		if appErr.next != nil {
			return GetRootError(appErr.next)
		} else {
			return appErr
		}
	} else {
		return err
	}
}

func (e *AppError) Format(s fmt.State, v rune) { xerrors.FormatError(e, s, v) }

func (e *AppError) FormatError(p xerrors.Printer) error {
	var message string
	if e.message != "" {
		message += fmt.Sprintf("%s", e.message)
	}

	p.Print(message)
	e.frame.Format(p)
	return e.next
}

func create(msg string) *AppError {
	var e AppError
	e.message = msg
	e.frame = xerrors.Caller(2)

	return &e
}

func New(msg string) *AppError {
	return create(msg)
}

func Errorf(format string, args ...interface{}) *AppError {
	return create(fmt.Sprintf(format, args...))
}

func Wrap(err error, msg ...string) *AppError {
	if err == nil {
		return nil
	}

	var m string
	if len(msg) != 0 {
		m = msg[0]
	}
	e := create(m)
	e.next = err
	return e
}

func (e *AppError) Unwrap() error { return e.next }

func Wrapf(err error, format string, args ...interface{}) *AppError {
	e := create(fmt.Sprintf(format, args...))
	e.next = err
	return e
}

func RootAs[T error](err error) (*T, bool) {
	rootErr := GetRootError(err)
	var target T
	ok := stdErrors.As(rootErr, &target)
	if !ok {
		return nil, false
	}
	return &target, ok
}

func AsAppError(err error) *AppError {
	if err == nil {
		return nil
	}

	var e *AppError
	if stdErrors.As(err, &e) {
		return e
	}
	return nil
}
