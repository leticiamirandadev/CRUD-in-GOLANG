package repository

import (
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/rest_err"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model"
)
type UserRepositoryInterface interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmail(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByID(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUser(
		userId string,
		userDomain model.UserDomainInterface,
	) *rest_err.RestErr
      DeleteUser(
		userId string,
	) *rest_err.RestErr

	FindUserByEmailAndPassword(
		email, password string,
	) (model.UserDomainInterface, *rest_err.RestErr)
}
