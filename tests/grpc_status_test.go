package errors

import (
	"testing"

	"github.com/jonh-dev/go-error/errors"
	"google.golang.org/grpc/codes"
)

func TestGRPCStatus(t *testing.T) {
	tests := []struct {
		name    string
		code    codes.Code
		message string
	}{
		{"InvalidArgument", codes.InvalidArgument, "Invalid argument"},
		{"DeadlineExceeded", codes.DeadlineExceeded, "Deadline exceeded"},
		{"NotFound", codes.NotFound, "Not found"},
		{"AlreadyExists", codes.AlreadyExists, "Already exists"},
		{"PermissionDenied", codes.PermissionDenied, "Permission denied"},
		{"ResourceExhausted", codes.ResourceExhausted, "Resource exhausted"},
		{"FailedPrecondition", codes.FailedPrecondition, "Failed precondition"},
		{"Aborted", codes.Aborted, "Aborted"},
		{"OutOfRange", codes.OutOfRange, "Out of range"},
		{"Unimplemented", codes.Unimplemented, "Unimplemented"},
		{"Internal", codes.Internal, "Internal error"},
		{"Unavailable", codes.Unavailable, "Unavailable"},
		{"DataLoss", codes.DataLoss, "Data loss"},
		{"Unauthenticated", codes.Unauthenticated, "Unauthenticated"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := errors.New(tt.code, tt.message)
			status := err.(*errors.Error).GRPCStatus()

			if status.Code() != tt.code {
				t.Errorf("Expected '%s', got '%s'", tt.code.String(), status.Code().String())
			}

			if status.Message() != tt.message {
				t.Errorf("Expected '%s', got '%s'", tt.message, status.Message())
			}
		})
	}
}
