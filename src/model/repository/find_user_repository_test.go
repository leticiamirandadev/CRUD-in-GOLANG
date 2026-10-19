package repository

import (
	"os"
	"testing"

	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model/repository/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_FindUserByEmail(t *testing.T) {
	collectionName := "user_collection_test"
	os.Setenv("MONGODB_USER_COLLECTION", collectionName)
	defer os.Clearenv()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("When_sending_a_valid_email_Returns_success", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB)
		oid := primitive.NewObjectID()

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmtNamespace(collectionName),
			mtest.FirstBatch,
			bson.D{
				{Key: "_id", Value: oid},
				{Key: "email", Value: "test@test.com"},
				{Key: "password", Value: "test"},
				{Key: "name", Value: "Test Name"},
				{Key: "age", Value: int8(20)},
			},
		))

		userDomain, err := repo.FindUserByEmail("test@test.com")

		assert.Nil(t, err)
		assert.NotNil(t, userDomain)
		assert.Equal(t, userDomain.GetID(), oid.Hex())
		assert.Equal(t, userDomain.GetEmail(), "test@test.com")
	})

	mt.Run("When_the_database_returns_no_documents_Returns_error", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB)

		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmtNamespace(collectionName),
			mtest.FirstBatch,
		))

		userDomain, err := repo.FindUserByEmail("test@test.com")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
		assert.Equal(t, err.Code, 404)
	})

	mt.Run("When_the_database_returns_error_Returns_error", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB)

		mt.AddMockResponses(primitive.D{{Key: "ok", Value: 0}})

		userDomain, err := repo.FindUserByEmail("test@test.com")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
		assert.Equal(t, err.Code, 500)
	})
}

func TestUserRepository_FindUserByID(t *testing.T) {
	collectionName := "user_collection_test"
	os.Setenv("MONGODB_USER_COLLECTION", collectionName)
	defer os.Clearenv()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("When_sending_a_valid_id_Returns_success", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB)
		oid := primitive.NewObjectID()

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmtNamespace(collectionName),
			mtest.FirstBatch,
			bson.D{
				{Key: "_id", Value: oid},
				{Key: "email", Value: "test@test.com"},
				{Key: "password", Value: "test"},
				{Key: "name", Value: "Test Name"},
				{Key: "age", Value: int8(20)},
			},
		))

		userDomain, err := repo.FindUserByID(oid.Hex())

		assert.Nil(t, err)
		assert.NotNil(t, userDomain)
		assert.Equal(t, userDomain.GetID(), oid.Hex())
	})

	mt.Run("When_the_database_returns_no_documents_Returns_error", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB)

		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmtNamespace(collectionName),
			mtest.FirstBatch,
		))

		userDomain, err := repo.FindUserByID(primitive.NewObjectID().Hex())

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
		assert.Equal(t, err.Code, 404)
	})
}

func TestUserRepository_FindUserByEmailAndPassword(t *testing.T) {
	collectionName := "user_collection_test"
	os.Setenv("MONGODB_USER_COLLECTION", collectionName)
	defer os.Clearenv()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("When_sending_a_valid_email_and_password_Returns_success", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB)
		oid := primitive.NewObjectID()

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmtNamespace(collectionName),
			mtest.FirstBatch,
			bson.D{
				{Key: "_id", Value: oid},
				{Key: "email", Value: "test@test.com"},
				{Key: "password", Value: "test"},
				{Key: "name", Value: "Test Name"},
				{Key: "age", Value: int8(20)},
			},
		))

		userDomain, err := repo.FindUserByEmailAndPassword("test@test.com", "test")

		assert.Nil(t, err)
		assert.NotNil(t, userDomain)
		assert.Equal(t, userDomain.GetID(), oid.Hex())
	})

	mt.Run("When_the_database_returns_no_documents_Returns_error", func(mt *mtest.T) {
		repo := NewUserRepository(mt.DB)

		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmtNamespace(collectionName),
			mtest.FirstBatch,
		))

		userDomain, err := repo.FindUserByEmailAndPassword("test@test.com", "test")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
		assert.Equal(t, err.Code, 403)
	})
}

func fmtNamespace(collection string) string {
	return "user_db_test." + collection
}

var _ = entity.UserEntity{}
