package repository

import (
	"context"
	"os"

	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/logger"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/rest_err"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model/repository/entity"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmailAndPassword(
	email, password string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmailAndPassword repository", zap.String("journey", "loginUser"))

	collectionName := os.Getenv("MONGODB_USER_COLLECTION")
	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}
	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "password", Value: password},
	}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "User or password is invalid"
			logger.Error(errorMessage, err, zap.String("journey", "loginUser"))
			return nil, rest_err.NewForbiddenError(errorMessage)
		}
		errorMessage := "Error trying to find user by email and password"
		logger.Error(errorMessage, err, zap.String("journey", "loginUser"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info(
		"FindUserByEmailAndPassword repository executed successfully",
		zap.String("journey", "loginUser"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID),
	)

	return converter.ConvertEntityToDomain(userEntity), nil
}
