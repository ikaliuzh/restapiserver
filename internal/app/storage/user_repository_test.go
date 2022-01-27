package storage_test

import (
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	st, teardown := storage.TestStorage(t, databaseURL)
	defer teardown("users")

	u, err := st.User().Create(&model.User{
		Email: "user@example.org",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	st, teardown := storage.TestStorage(t, databaseURL)
	defer teardown("users")

	email := "user@example.org"
	_, err := st.User().FindByEmail(email)
	assert.Error(t, err)

	st.User().Create(&model.User{
		Email: email,
	})

	u, err := st.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
