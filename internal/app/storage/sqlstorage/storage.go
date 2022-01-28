package sqlstorage

import (
	"database/sql"
	"http-rest-api/internal/app/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	db             *sql.DB
	userRepository *UserRepository
}

func New(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (st *Storage) User() storage.UserRepository {
	if st.userRepository != nil {
		return st.userRepository
	}
	st.userRepository = &UserRepository{
		storage: st,
	}
	return st.userRepository
}
