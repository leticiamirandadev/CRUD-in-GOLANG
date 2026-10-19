package service

import (
	"testing"

	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/rest_err"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserDomainService_FindUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockUserRepositoryInterface(ctrl)
	service := NewUserDomainService(repo)

	t.Run("When_sending_a_valid_id_Returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain("test@test.com", "test", "Test Name", 50)
		userDomain.SetID(id)

		repo.EXPECT().FindUserByID(id).Return(userDomain, nil)

		domain, err := service.FindUserByID(id)

		assert.Nil(t, err)
		assert.NotNil(t, domain)
		assert.Equal(t, domain.GetID(), id)
		assert.Equal(t, domain.GetEmail(), "test@test.com")
	})

	t.Run("When_the_repository_returns_error_Returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repo.EXPECT().FindUserByID(id).Return(nil, rest_err.NewBadRequest("user not found"))

		domain, err := service.FindUserByID(id)

		assert.NotNil(t, err)
		assert.Nil(t, domain)
	})
}

func TestUserDomainService_FindUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockUserRepositoryInterface(ctrl)
	service := NewUserDomainService(repo)

	t.Run("When_sending_a_valid_email_Returns_success", func(t *testing.T) {
		email := "test@test.com"
		userDomain := model.NewUserDomain(email, "test", "Test Name", 50)
		userDomain.SetID(primitive.NewObjectID().Hex())

		repo.EXPECT().FindUserByEmail(email).Return(userDomain, nil)

		domain, err := service.FindUserByEmail(email)

		assert.Nil(t, err)
		assert.NotNil(t, domain)
		assert.Equal(t, domain.GetEmail(), email)
	})

	t.Run("When_the_repository_returns_error_Returns_error", func(t *testing.T) {
		email := "error@test.com"

		repo.EXPECT().FindUserByEmail(email).Return(nil, rest_err.NewBadRequest("user not found"))

		domain, err := service.FindUserByEmail(email)

		assert.NotNil(t, err)
		assert.Nil(t, domain)
	})
}
