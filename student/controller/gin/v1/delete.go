package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/gin/util"
	"github.com/rebirthmonkey/go/pkg/log"
)

// Delete deletes an student by the student identifier.
// Only administrator can call this function.
func (u *controller) Delete(c *gin.Context) {
	log.L(c).Info("[GinServer] studentController: delete")

	if err := u.srv.NewStudentService().Delete(c.Param("name")); err != nil {
		util.WriteResponse(c, err, nil)

		return
	}

	var msg string = "deleted student " + c.Param("name")
	util.WriteResponse(c, nil, msg)
}