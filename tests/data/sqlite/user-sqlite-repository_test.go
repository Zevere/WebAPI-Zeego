package tests

import (
	"zeego/data/entities"
	"testing"
	"github.com/stretchr/testify/assert"
	"zeego/data/sqlite"
	"zeego/data"
	"os"
)

var repo data.UserRepository

func init() {
	if _, err := os.Stat("test.db"); os.IsNotExist(err) {
		return
	}
	if err := os.Remove("test.db"); err != nil {
		panic(err.Error())
	}
}

func TestUserRepo(t *testing.T) {
	var err data.Error
	repo, err = sqlite.NewUserSQLiteRepository("file:test.db?cache=shared&mode=memory")
	if err != nil {
		t.FailNow()
	}
	defer repo.Close()

	t.Run(
		"User Repository (SQLite)",
		func(t *testing.T) {
			t.Run("Insert new user Alice", testInsertNewUser)
			t.Run("Get user Alice by username", testGetUserAliceByName)
			t.Run("Try inserting invalid user", testInsertInvalidUser)
		},
	)
}

func testInsertNewUser(t *testing.T) {
	user := entities.User{
		Name:           "alice0",
		FirstName:      "Alice",
		PassphraseHash: "secret-passphrase",
	}

	err := repo.Insert(&user)

	assert.Nil(t, err)
}

func testGetUserAliceByName(t *testing.T) {
	user, err := repo.GetByName("ALICE0")

	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.Equal(t, "alice0", user.Name)
	assert.Equal(t, "Alice", user.FirstName)
	assert.Equal(t, "secret-passphrase", user.PassphraseHash)
	assert.NotEmpty(t, user.Id)
}

func testInsertInvalidUser(t *testing.T) {
	user := entities.User{
		Name:           "user1",
		PassphraseHash: "abcd",
		//FirstName:      "Left out",
	}

	err := repo.Insert(&user)

	assert.NotNil(t, err)
	assert.Equal(t, err.Code(), data.InvalidValue)
}
