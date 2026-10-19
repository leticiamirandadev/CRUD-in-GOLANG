package test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/controller"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model/repository"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model/service"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/test/connection"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	userController controller.UserControllerInterface
	db             *mongo.Database
	closeConn      func()
	mongoUserDB    string
)

func TestMain(m *testing.M) {
	mongoUserDB = "test_user_db"
	os.Setenv("MONGODB_USER_DB", mongoUserDB)
	os.Setenv("MONGODB_USER_COLLECTION", "user")

	db, closeConn = connection.OpenConnection()
	defer closeConn()

	repo := repository.NewUserRepository(db)
	serv := service.NewUserDomainService(repo)
	userController = controller.NewUserControllerInterface(serv)

	os.Exit(m.Run())
}

func TestIntegration_FindUserByEmail(t *testing.T) {
	ctx := context.Background()
	collectionName := os.Getenv("MONGODB_USER_COLLECTION")
	if collectionName == "" {
		collectionName = "user"
	}

	t.Run("When_user_does_not_exist_Returns_not_found", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ginCtx := CreateTestContext(recorder)

		params := gin.Params{
			{Key: "userEmail", Value: "nonexistent@test.com"},
		}
		MakeRequest(ginCtx, params, url.Values{}, "GET", nil)

		userController.FindUserByEmail(ginCtx)

		assert.Equal(t, 404, recorder.Code)
	})

	t.Run("When_user_exists_Returns_success", func(t *testing.T) {
		email := "test@test.com"
		_, err := db.Collection(collectionName).InsertOne(ctx, bson.M{
			"_id":      primitive.NewObjectID(),
			"email":    email,
			"password": "test123@",
			"name":     "Test Name",
			"age":      20,
		})
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()
		ginCtx := CreateTestContext(recorder)

		params := gin.Params{
			{Key: "userEmail", Value: email},
		}
		MakeRequest(ginCtx, params, url.Values{}, "GET", nil)

		userController.FindUserByEmail(ginCtx)

		assert.Equal(t, 200, recorder.Code)
	})
}

func CreateTestContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(recorder)
	return ctx
}

func MakeRequest(
	c *gin.Context,
	params gin.Params,
	urlValues url.Values,
	method string,
	body *bytes.Buffer,
) {
	var req *http.Request
	var err error

	if body == nil {
		req, err = http.NewRequest(method, "/?"+urlValues.Encode(), nil)
	} else {
		req, err = http.NewRequest(method, "/?"+urlValues.Encode(), body)
	}

	if err != nil {
		panic(err)
	}

	c.Request = req
	c.Params = params
}
