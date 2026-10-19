package model

import "github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/rest_err"

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetID() string

	EncryptPassword()
	GenerateToken() (string, *rest_err.RestErr)
	SetID(id string)
}
