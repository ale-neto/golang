package repository

import (
	"os"
	"testing"

	"github.com/ale-neto/golang/src/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestUserRepository_UpdateUser(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	err := os.Setenv("MONGODB_USER_DB", collectionName)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDb.Close()

	mtestDb.Run("when_sending_a_valid_user_return_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		domain := model.NewUserDomain(
			"test@test.com", "test", "test", 90)
		domain.SetID(bson.NewObjectID().Hex())

		err := repo.UpdateUser(domain.GetID(), domain)

		assert.Nil(t, err)
	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		domain := model.NewUserDomain(
			"test@test.com", "test", "test", 90)
		domain.SetID(bson.NewObjectID().Hex())
		err := repo.UpdateUser(domain.GetID(), domain)

		assert.NotNil(t, err)
	})
}
