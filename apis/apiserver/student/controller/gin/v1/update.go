package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/errcode"
	"github.com/rebirthmonkey/go/pkg/errors"
	"github.com/rebirthmonkey/go/pkg/gin/util"
	"github.com/rebirthmonkey/go/pkg/log"
	model "github.com/mushiguang/go/student/model/v1"
)

// Update updates a student's info by the student identifier.
func (u *controller) Update(c *gin.Context) {
	log.L(c).Info("[GinServer] studentController: update")

	var student model.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		log.L(c).Errorf("ErrBind: %s\n", err)
		util.WriteResponse(c, errors.WithCode(errcode.ErrBind, err.Error()), nil)

		return
	}

	student.Name = c.Param("name")

	if err := u.srv.NewStudentService().Update(&student); err != nil {
		util.WriteResponse(c, err, nil)

		return
	}

	util.WriteResponse(c, nil, student)
}