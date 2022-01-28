package teststorage_test

import (
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/storage/teststorage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	st := teststorage.New()

	u := model.TestUser(t)
	err := st.User().Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	st := teststorage.New()

	_, err := st.User().FindByEmail("user@example.org")
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = "user@example.org"
	st.User().Create(u)

	u, err = st.User().FindByEmail("user@example.org")
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
