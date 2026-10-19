package repository

import (
	"os"
	"testing"

	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model/repository/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_DeleteUser(t *testing.T) {
	collectionName := "user_collection_test"
	os.Setenv("MONGODB_USER_COLLECTION", collectionName)
	defer os.Clearenv()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("When_sending_a_valid_id_Returns_success", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB)
		
		mt.AddMockResponses(primitive.D{{Key: "ok", Value: 1}})

		err := repo.DeleteUser(primitive.NewObjectID().Hex())

		assert.Nil(t, err)
	})

	mt.Run("When_the_database_returns_error_Returns_error", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB)

		mt.AddMockResponses(primitive.D{{Key: "ok", Value: 0}})

		err := repo.DeleteUser(primitive.NewObjectID().Hex())

		assert.NotNil(t, err)
	})
}

var _ = entity.UserEntity{}
