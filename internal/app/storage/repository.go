package storage

import "http-rest-api/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}

type PreferencesRepository interface {
	Update(model.Preferences, int) error
	Get(userID int) (*model.Preferences, error)
}
