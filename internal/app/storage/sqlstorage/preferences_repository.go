package sqlstorage

import (
	"database/sql"
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/storage"
	"strings"
)

type PreferencesRepository struct {
	storage *Storage
}

// inserts new preferences or updates existing with new values
func (r *PreferencesRepository) Update(pref model.Preferences, userID int) error {
	var id *int
	return r.storage.db.QueryRow(
		"INSERT INTO preferences (user_id, tracked_tokens, fiat_currency) VALUES ($1, $2, $3) ON CONFLICT (user_id) DO UPDATE SET tracked_tokens = EXCLUDED.tracked_tokens, fiat_currency = EXCLUDED.fiat_currency  RETURNING (user_id)",
		userID,
		strings.Join(pref.TrackedTokens, " "),
		pref.FiatCurrency,
	).Scan(&id)
}

// finds preferences by user ID
func (r *PreferencesRepository) Get(userID int) (*model.Preferences, error) {
	pref := &model.Preferences{}
	var tmpTrackedTokens string

	if err := r.storage.db.QueryRow(
		"SELECT tracked_tokens, fiat_currency FROM preferences WHERE user_id = $1", userID,
	).Scan(&tmpTrackedTokens, &pref.FiatCurrency); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.ErrRecordNotFound
		}
		return nil, err
	}
	pref.TrackedTokens = strings.Split(tmpTrackedTokens, " ")
	return pref, nil
}
