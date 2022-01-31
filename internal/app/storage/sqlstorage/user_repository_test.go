package sqlstorage_test

import (
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/storage"
	"http-rest-api/internal/app/storage/sqlstorage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("users")

	st := sqlstorage.New(db)

	u := model.TestUser(t)
	err := st.User().Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("users")

	st := sqlstorage.New(db)

	_, err := st.User().FindByEmail("user@example.org")
	assert.EqualError(t, err, storage.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = "user@example.org"
	st.User().Create(u)

	u, err = st.User().FindByEmail("user@example.org")
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("users")

	st := sqlstorage.New(db)

	_, err := st.User().FindByEmail("user@example.org")
	assert.EqualError(t, err, storage.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = "user@example.org"
	st.User().Create(u)

	u, err = st.User().Find(u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
