package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/gin/util"
	"github.com/rebirthmonkey/go/pkg/log"
)

// List lists the students in the storage.
// Only administrator can call this function.
func (u *controller) List(c *gin.Context) {
	log.L(c).Info("[GinServer] studentController: list")

	students, err := u.srv.NewStudentService().List()
	if err != nil {
		util.WriteResponse(c, err, nil)

		return
	}

	util.WriteResponse(c, nil, students)
}