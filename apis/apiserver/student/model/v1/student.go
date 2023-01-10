package v1

import (
	"time"

	"github.com/rebirthmonkey/go/pkg/auth"
	"github.com/rebirthmonkey/go/pkg/metamodel"
	"github.com/rebirthmonkey/go/pkg/util"
	"gorm.io/gorm"
)

// Student represents a student restful resource. It is also used as data model.
type Student struct {
	metamodel.ObjectMeta `json:"metadata,omitempty"`

	Status      int64     `json:"status"              gorm:"column:status"    validate:"omitempty"`
	Nickname    string    `json:"nickname"            gorm:"column:nickname"  validate:"required,min=1,max=30"`
	Password    string    `json:"password,omitempty"  gorm:"column:password"  validate:"required"`
	Email       string    `json:"email"               gorm:"column:email"     validate:"required,email,min=1,max=100"`
	Phone       string    `json:"phone"               gorm:"column:phone"     validate:"omitempty"`
	IsAdmin     int64     `json:"isAdmin,omitempty"   gorm:"column:isAdmin"   validate:"omitempty"`
	TotalPolicy string    `json:"totalPolicy"         gorm:"-"                validate:"omitempty"`
	LoginedAt   time.Time `json:"loginedAt,omitempty" gorm:"column:loginedAt"`
}

// StudentList is the whole list of all students which have been stored in the storage.
type StudentList struct {
	// +optional
	metamodel.ListMeta `json:",inline"`

	Items []*Student `json:"items"`
}

// TableName maps to mysql table name.
func (u *Student) TableName() string {
	return "student"
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `Student` struct).
func (u *Student) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)

	return
}

// AfterCreate run after create database record.
func (u *Student) AfterCreate(tx *gorm.DB) error {
	u.InstanceID = util.GetInstanceID(u.ID, "student-")

	return tx.Save(u).Error
}