package controller

import (
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/rest_err"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/test/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserController_FindUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("When_ID_is_invalid_Returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := TestContext(recorder)
		params := gin.Params{
			{Key: "userId", Value: "invalid_id"},
		}
		MakeRequest(ctx, params, url.Values{}, "GET", nil)

		controller.FindUserByID(ctx)

		assert.Equal(t, 400, recorder.Code)
	})

	t.Run("When_service_returns_error_Returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := TestContext(recorder)
		id := primitive.NewObjectID().Hex()
		params := gin.Params{
			{Key: "userId", Value: id},
		}
		MakeRequest(ctx, params, url.Values{}, "GET", nil)

		service.EXPECT().FindUserByID(id).Return(nil, rest_err.NewInternalServerError("error trying to find user"))

		controller.FindUserByID(ctx)

		assert.Equal(t, 500, recorder.Code)
	})

	t.Run("When_user_is_found_successfully_Returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := TestContext(recorder)
		id := primitive.NewObjectID().Hex()
		params := gin.Params{
			{Key: "userId", Value: id},
		}
		MakeRequest(ctx, params, url.Values{}, "GET", nil)

		userDomain := model.NewUserDomain("test@test.com", "test", "Test Name", 50)
		userDomain.SetID(id)

		service.EXPECT().FindUserByID(id).Return(userDomain, nil)

		controller.FindUserByID(ctx)

		assert.Equal(t, 200, recorder.Code)
	})
}

func TestUserController_FindUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("When_service_returns_error_Returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := TestContext(recorder)
		params := gin.Params{
			{Key: "userEmail", Value: "test@test.com"},
		}
		MakeRequest(ctx, params, url.Values{}, "GET", nil)

		service.EXPECT().FindUserByEmail("test@test.com").Return(nil, rest_err.NewInternalServerError("error trying to find user"))

		controller.FindUserByEmail(ctx)

		assert.Equal(t, 500, recorder.Code)
	})

	t.Run("When_user_is_found_successfully_Returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := TestContext(recorder)
		params := gin.Params{
			{Key: "userEmail", Value: "test@test.com"},
		}
		MakeRequest(ctx, params, url.Values{}, "GET", nil)

		userDomain := model.NewUserDomain("test@test.com", "test", "Test Name", 50)
		userDomain.SetID(primitive.NewObjectID().Hex())

		service.EXPECT().FindUserByEmail("test@test.com").Return(userDomain, nil)

		controller.FindUserByEmail(ctx)

		assert.Equal(t, 200, recorder.Code)
	})
}
