package v1

import (
	model "github.com/mushiguang/go/student/model/v1"
	"github.com/mushiguang/go/student/repo"
	"time"

	"github.com/rebirthmonkey/go/pkg/metamodel"
	"golang.org/x/crypto/bcrypt"
)

// StudentService defines functions used to handle student request.
type StudentService interface {
	Create(student *model.Student) error
	Delete(studentname string) error
	Update(student *model.Student) error
	Get(studentname string) (*model.Student, error)
	List() (*model.StudentList, error)
}

// studentService is the StudentService instance to handle student request.
type studentService struct {
	repo repo.Repo
}

var _ StudentService = (*studentService)(nil)

// newStudentService creates and returns the student service instance.
func newStudentService(repo repo.Repo) StudentService {
	return &studentService{repo}
}

// Create creates a new student account.
func (u *studentService) Create(student *model.Student) error {
	hashedBytes, _ := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
	student.Password = string(hashedBytes)
	student.Status = 1
	student.LoginedAt = time.Now()

	return u.repo.StudentRepo().Create(student)
}

// Delete deletes the student by the student identifier.
func (u *studentService) Delete(studentname string) error {
	return u.repo.StudentRepo().Delete(studentname)
}

// Update updates a student account information.
func (u *studentService) Update(student *model.Student) error {
	updateStudent, err := u.Get(student.Name)
	if err != nil {
		return err
	}

	updateStudent.Nickname = student.Nickname
	updateStudent.Email = student.Email
	updateStudent.Phone = student.Phone
	updateStudent.Extend = student.Extend

	return u.repo.StudentRepo().Update(updateStudent)
}

// Get returns a student's info by the student identifier.
func (u *studentService) Get(studentname string) (*model.Student, error) {
	return u.repo.StudentRepo().Get(studentname)
}

// List returns all the related students.
func (u *studentService) List() (*model.StudentList, error) {
	students, err := u.repo.StudentRepo().List()
	if err != nil {
		return nil, err
	}

	infos := make([]*model.Student, 0)
	for _, student := range students.Items {
		infos = append(infos, &model.Student{
			ObjectMeta: metamodel.ObjectMeta{
				ID:        student.ID,
				Name:      student.Name,
				CreatedAt: student.CreatedAt,
				UpdatedAt: student.UpdatedAt,
			},
			Nickname: student.Nickname,
			Email:    student.Email,
			Phone:    student.Phone,
		})
	}

	return &model.StudentList{ListMeta: students.ListMeta, Items: infos}, nil
}