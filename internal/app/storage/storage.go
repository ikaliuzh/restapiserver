package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

func (st *Storage) Open() error {
	db, err := sql.Open("postgres", st.config.DatabaseURL)
	if err != nil {
		return nil
	}
	// as db is opened lazy we want to actually check that connection is established
	if err := db.Ping(); err != nil {
		return err
	}

	st.db = db

	return nil
}

func (st *Storage) Close() {
	st.db.Close()
}

func (st *Storage) User() *UserRepository {
	if st.userRepository != nil {
		return st.userRepository
	}
	st.userRepository = &UserRepository{
		storage: st,
	}
	return st.userRepository
}
