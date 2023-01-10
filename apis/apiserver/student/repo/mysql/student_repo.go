package mysql

import (
	"fmt"
	model "github.com/mushiguang/go/student/model/v1"
	studentRepoInterface "github.com/mushiguang/go/student/repo"
	"regexp"

	"github.com/rebirthmonkey/go/pkg/errcode"
	"github.com/rebirthmonkey/go/pkg/errors"
	"github.com/rebirthmonkey/go/pkg/log"
	"github.com/rebirthmonkey/go/pkg/mysql"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// studentRepo stores the student's info.
type studentRepo struct {
	dbEngine *gorm.DB
}

var _ studentRepoInterface.StudentRepo = (*studentRepo)(nil)

// newStudentRepo creates and returns a student storage.
func newStudentRepo(cfg *mysql.CompletedConfig) studentRepoInterface.StudentRepo {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		cfg.Studentname,
		cfg.Password,
		cfg.Host,
		cfg.Database,
		true,
		"Local")

	db, err := gorm.Open(mysqlDriver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Mysql connection fails %+v\n", err)
		return nil
	}

	return &studentRepo{dbEngine: db}
}

// close closes the repo's DB engine.
func (u *studentRepo) close() error {
	dbEngine, err := u.dbEngine.DB()
	if err != nil {
		return errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return dbEngine.Close()
}

// Create creates a new student account.
func (u *studentRepo) Create(student *model.Student) error {
	tmpStudent := model.Student{}
	u.dbEngine.Where("name = ?", student.Name).Find(&tmpStudent)
	if tmpStudent.Name != "" {
		err := errors.WithCode(errcode.ErrRecordAlreadyExist, "the created student already exit")

		log.Errorf("%+v", err)
		return err
	}

	err := u.dbEngine.Create(&student).Error
	if err != nil {
		if match, _ := regexp.MatchString("Duplicate entry", err.Error()); match {
			return errors.WrapC(err, errcode.ErrRecordAlreadyExist, "duplicate entry.")
		}

		return err
	}

	return nil
}

// Delete deletes the student by the student identifier.
func (u *studentRepo) Delete(studentname string) error {
	//tmpStudent := model.Student{}
	//u.dbEngine.Where("name = ?", studentname).Find(&tmpStudent)
	//if tmpStudent.Name == "" {
	//	err := errors.WithCode(errcode.ErrRecordNotFound, "the delete student not found")
	//	log.Errorf("%s\n", err)
	//	return err
	//}

	if err := u.dbEngine.Where("name = ?", studentname).Delete(&model.Student{}).Error; err != nil {
		return err
	}

	return nil
}

// Update updates a student account information.
func (u *studentRepo) Update(student *model.Student) error {
	tmpStudent := model.Student{}
	u.dbEngine.Where("name = ?", student.Name).Find(&tmpStudent)
	if tmpStudent.Name == "" {
		err := errors.WithCode(errcode.ErrRecordNotFound, "the update student not found")
		log.Errorf("%s\n", err)
		return err
	}

	if err := u.dbEngine.Save(student).Error; err != nil {
		return err
	}

	return nil
}

// Get returns a student's info by the student identifier.
func (u *studentRepo) Get(studentname string) (*model.Student, error) {
	student := &model.Student{}
	err := u.dbEngine.Where("name = ?", studentname).First(&student).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(errcode.ErrRecordNotFound, "the get student not found.")
		}

		return nil, errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return student, nil
}

// List returns all the related students.
func (u *studentRepo) List() (*model.StudentList, error) {
	ret := &model.StudentList{}

	d := u.dbEngine.
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)

	return ret, d.Error
}