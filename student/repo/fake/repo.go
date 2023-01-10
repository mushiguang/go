package fake

import (
	repo3 "github.com/mushiguang/go/student/repo"
	"sync"
)

// repo defines the APIServer storage.
type repo struct {
	studentRepo repo3.StudentRepo
}

var (
	r    repo3.Repo
	once sync.Once
)

var _ repo3.Repo = (*repo)(nil)

// Repo creates and returns the store client instance.
func Repo() (repo3.Repo, error) {
	once.Do(func() {
		r = repo{
			studentRepo: newStudentRepo(),
		}
	})

	return r, nil
}

// StudentRepo returns the student store client instance.
func (r repo) StudentRepo() repo3.StudentRepo {
	return r.studentRepo
}

// Close closes the repo.
func (r repo) Close() error {
	return nil
}