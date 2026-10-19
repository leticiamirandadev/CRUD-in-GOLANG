package service

import (
	"testing"

	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/rest_err"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/test/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockUserRepositoryInterface(ctrl)
	service := NewUserDomainService(repo)

	t.Run("When_sending_a_valid_id_Returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repo.EXPECT().DeleteUser(id).Return(nil)

		err := service.DeleteUser(id)

		assert.Nil(t, err)
	})

	t.Run("When_repository_returns_error_Returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repo.EXPECT().DeleteUser(id).Return(rest_err.NewInternalServerError("error trying to delete user"))

		err := service.DeleteUser(id)

		assert.NotNil(t, err)
		assert.Equal(t, err.Message, "error trying to delete user")
	})
}
