package repo

import (
	model "github.com/mushiguang/go/student/model/v1"
)

// StudentRepo defines the student resources.
type StudentRepo interface {
	Create(student *model.Student) error
	Delete(studentname string) error
	Update(student *model.Student) error
	Get(studentname string) (*model.Student, error)
	List() (*model.StudentList, error)
}