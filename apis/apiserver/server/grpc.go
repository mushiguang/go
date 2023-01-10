// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/rebirthmonkey/go/pkg/log"
	"github.com/mushiguang/go/apiserver/apis/apiserver/student/controller/grpc/v1"
	studentRepoMysql "github.com/mushiguang/go/apiserver/apis/apiserver/student/repo/mysql"
	"google.golang.org/grpc"
)

// InitGrpc initializes the Grpc server
func InitGrpc(server *grpc.Server) {
	log.Info("[GrpcServer] registry studentController")

	//studentRepoClient, err := studentRepoFake.Repo()
	//if err != nil {
	//	log.Fatalf("failed to create fake repo: %s", err.Error())
	//}

	studentRepoClient, err := studentRepoMysql.Repo(config.CompletedMysqlConfig)
	if err != nil {
		log.Fatalf("failed to create Mysql repo: %s", err.Error())
	}

	studentController := v1.NewController(studentRepoClient)
	v1.RegisterStudentServer(server, studentController)
}
