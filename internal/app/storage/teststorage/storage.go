package teststorage

import (
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/storage"
)

type Storage struct {
	userRepository *UserRepository
}

func New() *Storage {
	return &Storage{}
}

func (st *Storage) User() storage.UserRepository {
	if st.userRepository != nil {
		return st.userRepository
	}
	st.userRepository = &UserRepository{
		storage: st,
		users:   make(map[int]*model.User),
	}
	return st.userRepository
}
