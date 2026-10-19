package repository

import (
	"context"
	"os"

	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/logger"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)
func (ur *userRepository) DeleteUser(
	userId string,
) *rest_err.RestErr {
	logger.Info("Init deleteUser repository", zap.String("journey", "deleteUser"))

	collectionName := os.Getenv("MONGODB_USER_COLLECTION")
	if collectionName == "" {
		collectionName = "users"
	}
      collection := ur.databaseConnection.Collection(collectionName)

	userIdHex, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		logger.Error("Error trying to convert userId to ObjectID", err, zap.String("journey", "deleteUser"))
		return rest_err.NewBadRequest("Invalid user ID")
	}
      filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error trying to delete user", err, zap.String("journey", "deleteUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info(
		"DeleteUser repository executed successfully",
               zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
	)

	return nil
}
