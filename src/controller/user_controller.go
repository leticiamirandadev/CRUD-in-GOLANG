package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model/service"
)

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type userController struct {
	service service.UserDomainService
}
func NewUserControllerInterface(
	service service.UserDomainService,
) UserControllerInterface {
	return &userController{
		service: service,
	}
}
