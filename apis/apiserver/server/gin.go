// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/auth"
	"github.com/rebirthmonkey/go/pkg/gin/middleware"
	"github.com/rebirthmonkey/go/pkg/log"
	policyCtl "github.com/mushiguang/go/apiserver/apis/apiserver/policy/controller/gin/v1"
	policyRepo "github.com/mushiguang/go/apiserver/apis/apiserver/policy/repo"
	policyRepoMysql "github.com/mushiguang/go/apiserver/apis/apiserver/policy/repo/mysql"
	secretCtl "github.com/mushiguang/go/apiserver/apis/apiserver/secret/controller/gin/v1"
	secretRepo "github.com/mushiguang/go/apiserver/apis/apiserver/secret/repo"
	secretRepoMysql "github.com/mushiguang/go/apiserver/apis/apiserver/secret/repo/mysql"
	studentCtl "github.com/mushiguang/go/apiserver/apis/apiserver/student/controller/gin/v1"
	studentRepo "github.com/mushiguang/go/apiserver/apis/apiserver/student/repo"
	studentRepoMysql "github.com/mushiguang/go/apiserver/apis/apiserver/student/repo/mysql"
)

// InitGin initializes the Gin server
func InitGin(g *gin.Engine) {
	installRouterMiddleware(g)
	installController(g)
}

// installRouterMiddleware installs Gin server middlewares
func installRouterMiddleware(g *gin.Engine) {
	log.Info("[GinServer] registry LoggerMiddleware")
	g.Use(middleware.LoggerMiddleware())
}

// installController installs Gin handlers
func installController(g *gin.Engine) *gin.Engine {

	jwtStrategy, _ := newJWTAuth().(auth.JWTStrategy)
	g.POST("/login", jwtStrategy.LoginHandler)

	v1 := g.Group("/v1")
	{
		log.Info("[GinServer] registry studentHandler")
		studentv1 := v1.Group("/students")
		{
			//studentRepoClient, err := studentRepoFake.Repo()
			//if err != nil {
			//	log.Fatalf("failed to create fake repo: %s", err.Error())
			//}

			studentRepoClient, err := studentRepoMysql.Repo(config.CompletedMysqlConfig)
			if err != nil {
				log.Fatalf("failed to create Mysql repo: %s", err.Error())
			}
			studentRepo.SetClient(studentRepoClient)

			studentController := studentCtl.NewController(studentRepoClient)

			basicStrategy := newBasicAuth()
			studentv1.Use(basicStrategy.AuthFunc())

			studentv1.POST("", studentController.Create)
			studentv1.DELETE(":name", studentController.Delete)
			studentv1.PUT(":name", studentController.Update)
			studentv1.GET(":name", studentController.Get)
			studentv1.GET("", studentController.List)
		}

		log.Info("[GINServer] registry secretHandler")
		//secretv1 := v1.Group("/secrets")
		//secretv1.Use(jwtStrategy.AuthFunc())
		secretv1 := v1.Group("/secrets", jwtStrategy.AuthFunc())
		{
			secretRepoClient, _ := secretRepoMysql.Repo(config.CompletedMysqlConfig)
			secretRepo.SetClient(secretRepoClient)

			secretController := secretCtl.NewController(secretRepoClient)

			secretv1.POST("", secretController.Create)
			secretv1.DELETE(":name", secretController.Delete)
			secretv1.PUT(":name", secretController.Update)
			secretv1.GET(":name", secretController.Get)
			secretv1.GET("", secretController.List)
		}

		log.Info("[GINServer] registry policyHandler")
		policyv1 := v1.Group("/policies")
		{
			policyRepoClient, _ := policyRepoMysql.Repo(config.CompletedMysqlConfig)
			policyRepo.SetClient(policyRepoClient)

			policyController := policyCtl.NewController(policyRepoClient)

			policyv1.POST("", policyController.Create)
			policyv1.DELETE(":name", policyController.Delete)
			policyv1.PUT(":name", policyController.Update)
			policyv1.GET(":name", policyController.Get)
			policyv1.GET("", policyController.List)
		}
	}
	return g
}
