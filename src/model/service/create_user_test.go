package service

import (
	"testing"

	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/rest_err"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/test/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockUserRepositoryInterface(ctrl)
	service := NewUserDomainService(repo)

	t.Run("When_user_already_exists_Returns_error", func(t *testing.T) {
		userDomain := model.NewUserDomain("test@test.com", "test", "Test Name", 50)

		repo.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(userDomain, nil)

		domain, err := service.CreateUser(userDomain)

		assert.NotNil(t, err)
		assert.Nil(t, domain)
		assert.Equal(t, err.Message, "Email is already registered in another account")
	})

	t.Run("When_repository_returns_error_on_create_Returns_error", func(t *testing.T) {
		userDomain := model.NewUserDomain("test@test.com", "test", "Test Name", 50)

		repo.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)
		repo.EXPECT().CreateUser(gomock.Any()).Return(nil, rest_err.NewInternalServerError("error trying to create user"))

		domain, err := service.CreateUser(userDomain)

		assert.NotNil(t, err)
		assert.Nil(t, domain)
	})

	t.Run("When_user_is_created_successfully_Returns_success", func(t *testing.T) {
		userDomain := model.NewUserDomain("test@test.com", "test", "Test Name", 50)
		userDomain.SetID(primitive.NewObjectID().Hex())

		repo.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)
		repo.EXPECT().CreateUser(gomock.Any()).Return(userDomain, nil)

		domain, err := service.CreateUser(userDomain)

		assert.Nil(t, err)
		assert.NotNil(t, domain)
		assert.Equal(t, domain.GetEmail(), "test@test.com")
		assert.Equal(t, domain.GetName(), "Test Name")
	})
}
