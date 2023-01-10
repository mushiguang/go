package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/gin/util"
	"github.com/rebirthmonkey/go/pkg/log"
)

// Get gets a student's info by the student identifier.
func (u *controller) Get(c *gin.Context) {
	log.L(c).Info("[GinServer] studentController: get")

	student, err := u.srv.NewStudentService().Get(c.Param("name"))
	if err != nil {
		util.WriteResponse(c, err, nil)

		return
	}

	util.WriteResponse(c, nil, student)
}