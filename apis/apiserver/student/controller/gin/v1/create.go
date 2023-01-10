package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/errcode"
	"github.com/rebirthmonkey/go/pkg/errors"
	"github.com/rebirthmonkey/go/pkg/gin/util"
	"github.com/rebirthmonkey/go/pkg/log"
	model "github.com/mushiguang/go/student/model/v1"
)

// Create add new student to the storage.
func (u *controller) Create(c *gin.Context) {
	log.L(c).Info("[GinServer] studentController: create")

	var student model.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		log.L(c).Errorf("ErrBind: %s\n", err)
		util.WriteResponse(c, errors.WithCode(errcode.ErrBind, err.Error()), nil)

		return
	}

	if err := u.srv.NewStudentService().Create(&student); err != nil {
		util.WriteResponse(c, err, nil)
		return
	}

	util.WriteResponse(c, nil, student)
}