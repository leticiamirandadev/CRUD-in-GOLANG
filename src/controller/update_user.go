package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/logger"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/rest_err"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model/request"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userController) UpdateUser(c *gin.Context) {
	logger.Info("Init updateUser controller", zap.String("journey", "updateUser"))

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate userId", err, zap.String("journey", "updateUser"))
		errorMessage := rest_err.NewBadRequest("User id is not a valid hex")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}
      var userUpdateParams request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userUpdateParams); err != nil {
		logger.Error("Error trying to validate user update info", err, zap.String("journey", "updateUser"))
		restErr := rest_err.NewBadRequest("Some fields are incorrect")
		c.JSON(restErr.Code, restErr)
		return
	}
      name := ""
	if userUpdateParams.Name != nil {
		name = *userUpdateParams.Name
	}

	age := int8(0)
	if userUpdateParams.Age != nil {
		age = *userUpdateParams.Age
	}
      domain := model.NewUserUpdateDomain(name, age)

	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		logger.Error("Error trying to call updateUser service", err, zap.String("journey", "updateUser"))
		c.JSON(err.Code, err)
		return
	}
      logger.Info("UpdateUser controller executed successfully", zap.String("journey", "updateUser"))
	c.Status(http.StatusOK)
}
