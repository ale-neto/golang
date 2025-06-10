package repository

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_DeleteUser(t *testing.T) {
	// Define MOCK de client
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.CloneDatabase()

	// Simula a variável de ambiente usada no repo
	err := os.Setenv("MONGODB_USER_DB", "user_collection_test")
	assert.Nil(t, err)
	defer os.Clearenv()

	mt.Run("when_sending_a_valid_userId_return_success", func(mt *mtest.T) {
		// Mock de resposta para DeleteOne bem-sucedido: {ok:1, acknowledged:true, n:1}
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: int32(1)},
			{Key: "acknowledged", Value: true},
			{Key: "n", Value: int32(1)},
		})

		db := mt.Client.Database("user_database_test") // ou mt.DB()
		repo := NewUserRepository(db)

		err := repo.DeleteUser(context.Background(), "test-id")
		assert.Nil(t, err)
	})

	mt.Run("return_error_from_database_when_no_delete", func(mt *mtest.T) {
		// Mock simulando a falta de deleção: {ok:1, acknowledged:true, n:0}
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: int32(1)},
			{Key: "acknowledged", Value: true},
			{Key: "n", Value: int32(0)},
		})

		db := mt.Client.Database("user_database_test")
		repo := NewUserRepository(db)

		err := repo.DeleteUser(context.Background(), "test-id")
		assert.NotNil(t, err)
	})

	mt.Run("return_error_when_database_error", func(mt *mtest.T) {
		// Mock simulando erro no servidor: {ok:0}
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: int32(0)},
		})

		db := mt.Client.Database("user_database_test")
		repo := NewUserRepository(db)

		err := repo.DeleteUser(context.Background(), "test-id")
		assert.NotNil(t, err)
	})
}
