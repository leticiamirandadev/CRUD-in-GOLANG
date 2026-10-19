package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/logger"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)
func (uc *userController) DeleteUser(c *gin.Context) {
	logger.Info("Init deleteUser controller", zap.String("journey", "deleteUser"))

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate userId", err, zap.String("journey", "deleteUser"))
		errorMessage := rest_err.NewBadRequest("User id is not a valid hex")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}
      err := uc.service.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call deleteUser service", err, zap.String("journey", "deleteUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("DeleteUser controller executed successfully", zap.String("journey", "deleteUser"))
	c.Status(http.StatusOK)
}
