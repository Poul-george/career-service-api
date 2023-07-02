package errors

import "runtime"

type caller struct {
	file string
	line int
}

type appError struct {
	// 親のエラー
	parent  error
	message string
	caller  caller
}

func (e appError) Error() string {
	if e.parent != nil {
		return e.parent.Error()
	}
	return e.message
}

func New(m string) error {
	return newAppError(nil, m)
}

func Wrap(e error) error {
	if e == nil {
		return nil
	}
	return newAppError(e, "")
}

func newAppError(e error, m string) appError {
	_, file, line, _ := runtime.Caller(2)
	return appError{
		parent:  e,
		message: m,
		caller: caller{
			file: file,
			line: line,
		},
	}
}
