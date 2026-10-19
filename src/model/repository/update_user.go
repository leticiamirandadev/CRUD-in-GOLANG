package repository

import (
	"context"
	"os"

	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/logger"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/rest_err"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)
func (ur *userRepository) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init updateUser repository", zap.String("journey", "updateUser"))

	collectionName := os.Getenv("MONGODB_USER_COLLECTION")
	if collectionName == "" {
		collectionName = "users"
	}
      collection := ur.databaseConnection.Collection(collectionName)

	value := converter.ConvertDomainToEntity(userDomain)

	userIdHex, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		logger.Error("Error trying to convert userId to ObjectID", err, zap.String("journey", "updateUser"))
		return rest_err.NewBadRequest("Invalid user ID")
	}
      filter := bson.D{{Key: "_id", Value: userIdHex}}
	update := bson.D{{Key: "$set", Value: value}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Error trying to update user", err, zap.String("journey", "updateUser"))
		return rest_err.NewInternalServerError(err.Error())
	}
      logger.Info(
		"UpdateUser repository executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)

	return nil
}
