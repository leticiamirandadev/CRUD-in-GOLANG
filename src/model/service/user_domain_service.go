package service

import (
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/logger"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/rest_err"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model/repository"
	"go.uber.org/zap"
)

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(string) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
	LoginUserServices(model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr)
}
type userDomainService struct {
	userRepository repository.UserRepositoryInterface
}

func NewUserDomainService(
	userRepository repository.UserRepositoryInterface,
) UserDomainService {
	return &userDomainService{
		userRepository: userRepository,
	}
}
func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser service", zap.String("journey", "createUser"))

	existingUser, _ := ud.FindUserByEmail(userDomain.GetEmail())
	if existingUser != nil {
		return nil, rest_err.NewBadRequest("Email is already registered in another account")
	}
      userDomain.EncryptPassword()

	userDomainResult, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info(
		"CreateUser service executed successfully",
		zap.String("userId", userDomainResult.GetID()),
		zap.String("journey", "createUser"),
	)
      return userDomainResult, nil
}

func (ud *userDomainService) FindUserByID(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID service", zap.String("journey", "findUserByID"))
	return ud.userRepository.FindUserByID(id)
}
func (ud *userDomainService) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail service", zap.String("journey", "findUserByEmail"))
	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init updateUser service", zap.String("journey", "updateUser"))
	return ud.userRepository.UpdateUser(userId, userDomain)
}
func (ud *userDomainService) DeleteUser(
	userId string,
) *rest_err.RestErr {
	logger.Info("Init deleteUser service", zap.String("journey", "deleteUser"))
	return ud.userRepository.DeleteUser(userId)
}

func (ud *userDomainService) LoginUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, string, *rest_err.RestErr) {
	logger.Info("Init loginUser services", zap.String("journey", "loginUser"))

	userDomain.EncryptPassword()
      user, err := ud.userRepository.FindUserByEmailAndPassword(
		userDomain.GetEmail(),
		userDomain.GetPassword(),
	)

	if err != nil {
		return nil, "", err
	}

	token, errToken := user.GenerateToken()
	if errToken != nil {
		return nil, "", errToken
	}
      return user, token, nil
}
