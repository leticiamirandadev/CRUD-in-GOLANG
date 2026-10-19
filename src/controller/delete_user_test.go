package controller

import (
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/rest_err"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/test/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserController_DeleteUser(t *testing.T) {
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
		MakeRequest(ctx, params, url.Values{}, "DELETE", nil)

		controller.DeleteUser(ctx)

		assert.Equal(t, 400, recorder.Code)
	})

	t.Run("When_service_returns_error_Returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := TestContext(recorder)
		id := primitive.NewObjectID().Hex()
		params := gin.Params{
			{Key: "userId", Value: id},
		}
		MakeRequest(ctx, params, url.Values{}, "DELETE", nil)

		service.EXPECT().DeleteUser(id).Return(rest_err.NewInternalServerError("error trying to delete user"))

		controller.DeleteUser(ctx)

		assert.Equal(t, 500, recorder.Code)
	})

	t.Run("When_user_is_deleted_successfully_Returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := TestContext(recorder)
		id := primitive.NewObjectID().Hex()
		params := gin.Params{
			{Key: "userId", Value: id},
		}
		MakeRequest(ctx, params, url.Values{}, "DELETE", nil)

		service.EXPECT().DeleteUser(id).Return(nil)

		controller.DeleteUser(ctx)

		assert.Equal(t, 200, recorder.Code)
	})
}
