package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mushiguang/go/student/repo"
	srv "github.com/mushiguang/go/student/service/v1"
)

// Controller creates a student handler interface for student resource.
type Controller interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
	List(c *gin.Context)
}

// controller creates a student handler used to handle request for student resource.
type controller struct {
	srv srv.Service
}

// NewController creates a student handler.
func NewController(repo repo.Repo) Controller {
	return &controller{
		srv: srv.NewService(repo),
	}
}