package teststorage

import (
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/storage"
)

type PreferencesRepository struct {
	storage    *Storage
	preferenes map[int]*model.Preferences
}

func (r *PreferencesRepository) Update(pref model.Preferences, userID int) error {
	r.preferenes[userID] = &pref
	return nil
}

func (r *PreferencesRepository) Get(userID int) (*model.Preferences, error) {
	pref, ok := r.preferenes[userID]
	if !ok {
		return nil, storage.ErrRecordNotFound
	}
	return pref, nil
}
