package errors

import (
	"google.golang.org/grpc/status"
)

/*
GRPCStatus retorna o status do erro.

- e: erro.

@params e error

@returns *status.Status
*/
func (e *Error) GRPCStatus() *status.Status {
	return status.New(e.Code, e.Message)
}
