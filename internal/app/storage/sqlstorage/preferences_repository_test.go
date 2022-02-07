package sqlstorage_test

import (
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/storage"
	"http-rest-api/internal/app/storage/sqlstorage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreferencesRepository_Update(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("users", "preferences")

	st := sqlstorage.New(db)

	u := model.TestUser(t)
	st.User().Create(u)

	pref := model.TestPreferences(t)

	// insert new preferences
	err := st.Preferences().Update(pref, u.ID)
	assert.NoError(t, err)

	// check that preferences were correctly inserted
	rpref, err := st.Preferences().Get(u.ID)

	assert.NoError(t, err)
	assert.Equal(t, rpref.FiatCurrency, pref.FiatCurrency)
	assert.Equal(t, rpref.TrackedTokens, pref.TrackedTokens)

	// try to update existing preferences
	newpref := model.TestPreferences(t)
	newpref.FiatCurrency = "UAH"

	err = st.Preferences().Update(newpref, u.ID)
	assert.NoError(t, err)

	rpref, err = st.Preferences().Get(u.ID)

	assert.NoError(t, err)
	assert.Equal(t, rpref.FiatCurrency, newpref.FiatCurrency)
	assert.Equal(t, rpref.TrackedTokens, newpref.TrackedTokens)
}

func TestPreferencesRepository_Get(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("users", "preferences")

	st := sqlstorage.New(db)

	_, err := st.Preferences().Get(42)

	assert.EqualError(t, err, storage.ErrRecordNotFound.Error())
}
