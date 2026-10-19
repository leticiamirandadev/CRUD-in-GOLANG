package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/rest_err"
)

type userLoginDomain struct {
	email    string
	password string
}
func NewUserLoginDomain(
	email, password string,
) UserDomainInterface {
	return &userLoginDomain{
		email:    email,
		password: password,
	}
}
func (ud *userLoginDomain) EncryptPassword() {
	hasher := md5.New()
	hasher.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hasher.Sum(nil))
}

func (ud *userLoginDomain) GenerateToken() (string, *rest_err.RestErr) {
	return "", rest_err.NewInternalServerError("Not implemented for login domain")
}
func (ud *userLoginDomain) GetEmail() string     { return ud.email }
func (ud *userLoginDomain) GetPassword() string { return ud.password }
func (ud *userLoginDomain) GetName() string     { return "" }
func (ud *userLoginDomain) GetAge() int8        { return 0 }
func (ud *userLoginDomain) GetID() string       { return "" }
func (ud *userLoginDomain) SetID(id string)     {}
