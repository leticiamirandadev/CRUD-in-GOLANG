package repository

import (
	"os"
	"testing"

	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_CreateUser(t *testing.T) {
	collectionName := "user_collection_test"

	os.Setenv("MONGODB_USER_COLLECTION", collectionName)
	defer os.Clearenv()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
      mt.Run("When_sending_a_valid_domain_Returns_success", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB)

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		userDomain := model.NewUserDomain(
			"test@test.com",
			"test",
			"Test Name",
			20,
		)
              userDomainResult, err := repo.CreateUser(userDomain)

		assert.Nil(t, err)
		assert.NotNil(t, userDomainResult)
		assert.NotEmpty(t, userDomainResult.GetID())
		assert.Equal(t, userDomainResult.GetEmail(), "test@test.com")
		assert.Equal(t, userDomainResult.GetName(), "Test Name")
		assert.Equal(t, userDomainResult.GetAge(), int8(20))
	})
      mt.Run("When_the_database_returns_error_Returns_error", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB)

		mt.AddMockResponses(primitive.D{{Key: "ok", Value: 0}})

		userDomain := model.NewUserDomain(
			"test@test.com",
			"test",
			"Test Name",
			20,
		)
              userDomainResult, err := repo.CreateUser(userDomain)

		assert.NotNil(t, err)
		assert.Nil(t, userDomainResult)
		assert.Equal(t, err.Code, 500)
	})
}
