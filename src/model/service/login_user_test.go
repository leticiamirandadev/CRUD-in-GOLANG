package service

import (
	"os"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/rest_err"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/test/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_LoginUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockUserRepositoryInterface(ctrl)
	service := NewUserDomainService(repo)

	t.Run("When_repository_returns_error_Returns_error", func(t *testing.T) {
		userDomain := model.NewUserDomain("test@test.com", "test", "Test Name", 50)

		repo.EXPECT().FindUserByEmailAndPassword(gomock.Any(), gomock.Any()).Return(nil, rest_err.NewInternalServerError("error trying to find user"))

		domain, token, err := service.LoginUserServices(userDomain)

		assert.NotNil(t, err)
		assert.Nil(t, domain)
		assert.Equal(t, token, "")
	})

	t.Run("When_user_is_logged_successfully_Returns_success", func(t *testing.T) {
		os.Setenv("JWT_SECRET", "test_secret")
		defer os.Clearenv()

		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain("test@test.com", "test", "Test Name", 50)
		userDomain.SetID(id)

		repo.EXPECT().FindUserByEmailAndPassword(gomock.Any(), gomock.Any()).Return(userDomain, nil)

		domain, token, err := service.LoginUserServices(userDomain)

		assert.Nil(t, err)
		assert.NotNil(t, domain)
		assert.NotEqual(t, token, "")

		// Validar o token JWT gerado
		tokenParsed, errJwt := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte("test_secret"), nil
		})

		assert.Nil(t, errJwt)
		assert.True(t, tokenParsed.Valid)
	})
}
