package user

import (
	"database/sql"
	"fmt"

	"github.com/MohamedAklamaash/GO_Simple_CRUD_Backend_WIth_SQL/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email=?", email)
	if err != nil {
		return nil, err
	}
	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.Id == "" {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(&user.Id, &user.CreatedAt, &user.FirstName, &user.LastName, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Store) GetUserById(id string) (*types.User, error) {
	return nil, nil
}

func (s *Store) CreateUser(u types.User) error {
	return nil
}
