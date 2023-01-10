// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/mushiguang/go/apiserver/apis/apiserver/secret/repo"
)

// Service defines functions used to return resource interface.
type Service interface {
	NewSecretService() SecretService
}

// service is the business logic of the student resource handling.
type service struct {
	repo repo.Repo
}

var _ Service = (*service)(nil)

// NewService returns service instance of the Service interface.
func NewService(repo repo.Repo) Service {
	return &service{repo}
}

// NewSecretService returns a secret service instance.
func (s *service) NewSecretService() SecretService {
	return newSecretService(s.repo)
}
