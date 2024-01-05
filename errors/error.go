package errors

import (
	"google.golang.org/grpc/codes"
)

type Error struct {
	Code    codes.Code
	Message string
}

/*
New retorna um novo erro com o código e a mensagem informados.

- code: código do erro.

- message: mensagem do erro.

@params code codes.Code

@params message string

@returns error
*/
func New(code codes.Code, message string) error {
	return &Error{Code: code, Message: message}
}

/*
Error retorna a mensagem do erro.

- e: erro.

@params e error

@returns string
*/
func (e *Error) Error() string {
	return e.Message
}
