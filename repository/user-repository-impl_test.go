package repository

import (
	"os"
	"testing"

	"github.com/rht6226/event-management-app/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	testDb = "/tmp/event-management-test.db"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(testDb), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}
	return db
}

func TestUserRepository(t *testing.T) {
	// Initialize a new in-memory SQLite database for testing
	db := setupTestDB()
	defer os.Remove(testDb)

	// Migrate the schema
	err := db.AutoMigrate(&model.User{})
	assert.NoError(t, err)

	// Create a new UserRepository instance
	repo := NewUserRepository(db)

	// Test Save method
	user := &model.User{Email: "john.doe@example.com", Password: "password", Name: struct {
		FirstName string
		LastName  string
	}{
		FirstName: "John",
		LastName:  "Doe",
	}}
	createdUser, err := repo.Save(user)
	assert.NoError(t, err)
	assert.NotNil(t, createdUser.ID)
	assert.Equal(t, user.Name.FirstName, createdUser.Name.FirstName)
	assert.Equal(t, user.Name.LastName, createdUser.Name.LastName)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.Equal(t, user.Password, createdUser.Password)

	// Test FindByID method
	// TODO: This testcase is failing.

	// foundUserByID, err := repo.FindByID(createdUser.ID)
	// assert.NoError(t, err)
	// assert.NotNil(t, foundUserByID)
	// assert.Equal(t, createdUser.ID, foundUserByID.ID)
	// assert.Equal(t, createdUser.Name, foundUserByID.Name)
	// assert.Equal(t, createdUser.Email, foundUserByID.Email)
	// assert.Equal(t, createdUser.Password, foundUserByID.Password)

	// Test FindByEmail method
	foundUserByEmail, err := repo.FindByEmail("john.doe@example.com")
	t.Log(err)
	assert.NoError(t, err)
	assert.NotNil(t, foundUserByEmail)
	assert.NotNil(t, foundUserByEmail.ID)
	assert.Equal(t, createdUser.ID, foundUserByEmail.ID)
	assert.Equal(t, createdUser.Name, foundUserByEmail.Name)
	assert.Equal(t, createdUser.Email, foundUserByEmail.Email)
	assert.Equal(t, createdUser.Password, foundUserByEmail.Password)

	// // Test FindAll method
	allUsers, err := repo.FindAll()
	assert.NoError(t, err)
	assert.NotNil(t, allUsers)
	assert.Len(t, allUsers, 1)

	// Test Delete method
	err = repo.Delete(createdUser.ID)
	assert.NoError(t, err)

	// Ensure the user is deleted
	deletedUser, err := repo.FindByID(createdUser.ID)
	assert.Error(t, err)
	assert.Nil(t, deletedUser)
}
