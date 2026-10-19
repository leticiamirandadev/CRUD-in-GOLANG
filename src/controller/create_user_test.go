package controller

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/rest_err"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/controller/model/request"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/test/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserController_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("When_request_is_invalid_Returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := TestContext(recorder)

		userRequest := request.UserRequest{
			Email:    "invalid_email",
			Password: "123",
			Name:     "Test",
			Age:      10,
		}
		jsonValue, _ := json.Marshal(userRequest)
		body := bytes.NewBuffer(jsonValue)

		MakeRequest(ctx, gin.Params{}, url.Values{}, "POST", body)

		controller.CreateUser(ctx)

		assert.Equal(t, 400, recorder.Code)
	})

	t.Run("When_service_returns_error_Returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := TestContext(recorder)

		userRequest := request.UserRequest{
			Email:    "test@test.com",
			Password: "test123@",
			Name:     "Test Name",
			Age:      20,
		}
		jsonValue, _ := json.Marshal(userRequest)
		body := bytes.NewBuffer(jsonValue)

		MakeRequest(ctx, gin.Params{}, url.Values{}, "POST", body)

		service.EXPECT().CreateUser(gomock.Any()).Return(nil, rest_err.NewInternalServerError("error trying to create user"))

		controller.CreateUser(ctx)

		assert.Equal(t, 500, recorder.Code)
	})

	t.Run("When_user_is_created_successfully_Returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := TestContext(recorder)

		userRequest := request.UserRequest{
			Email:    "test@test.com",
			Password: "test123@",
			Name:     "Test Name",
			Age:      20,
		}
		jsonValue, _ := json.Marshal(userRequest)
		body := bytes.NewBuffer(jsonValue)

		MakeRequest(ctx, gin.Params{}, url.Values{}, "POST", body)

		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain("test@test.com", "test123@", "Test Name", 20)
		userDomain.SetID(id)

		service.EXPECT().CreateUser(gomock.Any()).Return(userDomain, nil)

		controller.CreateUser(ctx)

		assert.Equal(t, 200, recorder.Code)
	})
}
