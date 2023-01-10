package v1

import (
	"context"
	"github.com/mushiguang/go/apis/apiserver/student/repo"
	srv "github.com/mushiguang/go/apis/apiserver/student/service/v1"
)

// Controller creates a GRPC student interface for student resource.
type Controller interface {
	ListStudents(ctx context.Context, r *ListStudentsRequest) (*ListStudentsResponse, error)
}

// controller creates a GRPC student handler used to handle request for student resource.
type controller struct {
	srv srv.Service
	UnimplementedStudentServer
}

// NewController creates a GRPC student handler.
func NewController(repo repo.Repo) *controller {
	return &controller{
		srv: srv.NewService(repo),
	}
}