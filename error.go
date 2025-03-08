package inerr

import (
	"fmt"
)

func New(format string, args ...any) error {
	return fmt.Errorf(format, args...)
}

type msgError struct {
	err error
	msg string
}

func (s *msgError) Message() string {
	return s.msg
}

func (s *msgError) Error() string {
	return fmt.Sprintf("%s: %v", s.msg, s.err)
}

func (s *msgError) Unwrap() error {
	return s.err
}

func Wrap(err error, format string, args ...any) error {
	return &msgError{
		err: err,
		msg: fmt.Sprintf(format, args...),
	}
}
