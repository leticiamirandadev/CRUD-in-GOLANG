package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/logger"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/rest_err"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/controller/model/request"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/view"
	"go.uber.org/zap"
)

func (uc *userController) LoginUser(c *gin.Context) {
	logger.Info("Init loginUser controller", zap.String("journey", "loginUser"))

	var userLoginRequest request.UserLoginRequest
      if err := c.ShouldBindJSON(&userLoginRequest); err != nil {
		logger.Error("Error trying to validate user login request data", err, zap.String("journey", "loginUser"))
		restErr := rest_err.NewBadRequest("Some fields are incorrect")
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserLoginDomain(
		userLoginRequest.Email,
		userLoginRequest.Password,
	)

	domainResult, token, err := uc.service.LoginUserServices(domain)
	if err != nil {
		logger.Error("Error trying to call loginUser service", err, zap.String("journey", "loginUser"))
		c.JSON(err.Code, err)
		return
	}

	c.Header("Authorization", token)
      logger.Info("LoginUser controller executed successfully", zap.String("journey", "loginUser"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
