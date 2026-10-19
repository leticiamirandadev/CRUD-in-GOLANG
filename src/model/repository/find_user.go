package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/logger"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/rest_err"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model/repository/entity"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail repository", zap.String("journey", "findUserByEmail"))

	collectionName := os.Getenv("MONGODB_USER_COLLECTION")
	if collectionName == "" {
		collectionName = "users"
	}

	collection := ur.databaseConnection.Collection(collectionName)
      userEntity := &entity.UserEntity{}
	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this email: %s", email)
			logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
              errorMessage := fmt.Sprintf("Error trying to find user by email: %s", err.Error())
		logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info(
		"FindUserByEmail repository executed successfully",
		zap.String("userId", userEntity.ID),
		zap.String("journey", "findUserByEmail"),
	)
      return converter.ConvertEntityToDomain(userEntity), nil
}

func (ur *userRepository) FindUserByID(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID repository", zap.String("journey", "findUserByID"))

	collectionName := os.Getenv("MONGODB_USER_COLLECTION")
	if collectionName == "" {
		collectionName = "users"
	}
      collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("Error trying to convert id to ObjectID", err, zap.String("journey", "findUserByID"))
		return nil, rest_err.NewBadRequest("Invalid user ID")
	}
      filter := bson.D{{Key: "_id", Value: objID}}
	err = collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this ID: %s", id)
			logger.Error(errorMessage, err, zap.String("journey", "findUserByID"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
              errorMessage := fmt.Sprintf("Error trying to find user by ID: %s", err.Error())
		logger.Error(errorMessage, err, zap.String("journey", "findUserByID"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info(
		"FindUserByID repository executed successfully",
		zap.String("userId", userEntity.ID),
		zap.String("journey", "findUserByID"),
	)
      return converter.ConvertEntityToDomain(userEntity), nil
}
