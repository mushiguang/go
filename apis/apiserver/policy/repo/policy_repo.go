// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package repo

import (
	model "github.com/mushiguang/go/apis/apiserver/policy/model/v1"
)

// PolicyRepo defines the secret resources.
type PolicyRepo interface {
	Create(policy *model.Policy) error
	Delete(studentname string, policyName string) error
	DeleteByStudent(studentname string) error
	Update(policy *model.Policy) error
	Get(studentname string, policyName string) (*model.Policy, error)
	List(studentname string) (*model.PolicyList, error)
}
