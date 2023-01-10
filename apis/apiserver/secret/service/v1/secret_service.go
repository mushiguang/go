// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	model "github.com/mushiguang/go/apiserver/apis/apiserver/secret/model/v1"
	"github.com/mushiguang/go/apiserver/apis/apiserver/secret/repo"
)

type SecretService interface {
	Create(secret *model.Secret) error
	Update(secret *model.Secret) error
	Delete(studentname, secretID string) error
	Get(studentname, secretID string) (*model.Secret, error)
	List(studentname string) (*model.SecretList, error)
}

type secretService struct {
	repo repo.Repo
}

var _ SecretService = (*secretService)(nil)

// newSecretService creates and returns the secret service instance.
func newSecretService(repo repo.Repo) SecretService {
	return &secretService{repo: repo}
}

func (s *secretService) Create(secret *model.Secret) error {
	return s.repo.SecretRepo().Create(secret)
}

func (s *secretService) Delete(studentname, secretName string) error {
	return s.repo.SecretRepo().Delete(studentname, secretName)
}

func (s *secretService) Update(secret *model.Secret) error {
	return s.repo.SecretRepo().Update(secret)
}

func (s *secretService) Get(studentname, secretName string) (*model.Secret, error) {
	return s.repo.SecretRepo().Get(studentname, secretName)
}

func (s *secretService) List(studentname string) (*model.SecretList, error) {
	return s.repo.SecretRepo().List(studentname)
}
