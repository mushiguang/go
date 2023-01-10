package fake

import (
	"fmt"
	"github.com/rebirthmonkey/go/pkg/errcode"
	"github.com/rebirthmonkey/go/pkg/errors"
	"github.com/rebirthmonkey/go/pkg/metamodel"
	model "github.com/mushiguang/go/student/model/v1"
	StudentRepoInterface "github.com/mushiguang/go/student/repo"
)

// studentRepo stores the student's info.
type StudentRepo struct {
	dbEngine []*model.Student
}

var _ studentRepoInterface.StudentRepo = (*studentRepo)(nil)

// newUserRepo creates and returns a student storage.
func newStudentRepo() studentRepoInterface.StudentRepo {

	students := make([]*model.Student, 0)
	for i := 1; i <= 10; i++ {
		students = append(students, &model.Student{
			ObjectMeta: metamodel.ObjectMeta{
				Name: fmt.Sprintf("student%d", i),
				ID:   uint64(i),
			},
			Nickname: fmt.Sprintf("student%d", i),
			Password: fmt.Sprintf("Student%d@2022", i),
			Email:    fmt.Sprintf("student%d@qq.com", i),
		})
	}

	return &studentRepo{
		dbEngine: students,
	}
}

// Create creates a new student account.
func (u *studentRepo) Create(student *model.Student) error {
	for _, u := range u.dbEngine {
		if u.Name == student.Name {
			return errors.WithCode(errcode.ErrRecordAlreadyExist, "record already exist")
		}
	}

	if len(u.dbEngine) > 0 {
		student.ID = u.dbEngine[len(u.dbEngine)-1].ID + 1
	}
	u.dbEngine = append(u.dbEngine, student)

	return nil
}

// Delete deletes the student by the student identifier.
func (u *studentRepo) Delete(studentname string) error {
	newStudents := make([]*model.Student, 0)

	for i := 0; i < len(u.dbEngine); i++ {
		if u.dbEngine[i].Name == studentname {
			newStudents = append(u.dbEngine[:i], u.dbEngine[i+1:]...)
			break
		}
	}

	if len(newStudents) == 0 {
		return errors.WithCode(errcode.ErrRecordNotFound, "record not found")
	}

	u.dbEngine = newStudents
	return nil
}

// Update updates a student account information.
func (u *studentRepo) Update(student *model.Student) error {
	if err := u.Delete(student.Name); err != nil {
		return err
	}

	return u.Create(student)
}

// Get returns a student's info by the student identifier.
func (u *studentRepo) Get(studentname string) (*model.Student, error) {
	for _, u := range u.dbEngine {
		if u.Name == studentname {
			return u, nil
		}
	}

	return nil, errors.WithCode(errcode.ErrRecordNotFound, "record not found")
}

// List returns all the related students.
func (u *studentRepo) List() (*model.StudentList, error) {
	students := make([]*model.Student, 0)
	i := 0
	for _, student := range u.dbEngine {
		students = append(students, student)
		i++
	}

	return &model.StudentList{
		ListMeta: metamodel.ListMeta{
			TotalCount: int64(len(u.dbEngine)),
		},
		Items: students,
	}, nil
}