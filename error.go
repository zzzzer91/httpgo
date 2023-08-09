package httpgo

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/zzzzer91/gopkg/stackx"
)

type StatusError struct {
	statusCode int
	msg        string
	stack      *stackx.Stack
}

func NewStatusError(statusCode int, msg string) *StatusError {
	return &StatusError{
		statusCode: statusCode,
		msg:        msg,
		stack:      stackx.Callers(1),
	}
}

func (e *StatusError) StatusCode() int {
	return e.statusCode
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("statusCode: %d, msg: %s", e.statusCode, e.msg)
}

func (e *StatusError) StackTrace() errors.StackTrace {
	return e.stack.StackTrace()
}
