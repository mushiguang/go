package v1

import (
	"github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/student/repo"
)

// Service defines functions used to return resource interface.
type Service interface {
	NewStudentService() StudentService
}

// service is the business logic of the student resource handling.
type service struct {
	repo repo.Repo
}

var _ Service = (*service)(nil)

// NewService returns service instance of the Service interface.
func NewService(repo repo.Repo) Service {
	return &service{repo}
}

// NewStudentService returns a student service instance.
func (s *service) NewStudentService() StudentService {
	return newStudentService(s.repo)
}