package repo

import (
	"database/sql"

	"project-name/internal/models"
)

type UserRepo interface {
	EmailExists(email string) (bool, error)

	Add(req *models.User) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetById(email string) (*models.User, error)
}

type userSql struct {
	conn *sql.DB
}

func (u *userSql) EmailExists(email string) (bool, error) {
	var userId string

	query := `SELECT id FROM users WHERE email = $1;`

	err := u.conn.QueryRow(query, email).Scan(&userId)

	if err != nil {
		if err == sql.ErrNoRows {
			// Email does not exist
			return false, nil
		}
		// An error occurred while executing the query
		return true, err
	}

	// Email already exists
	return true, nil
}

func (m *userSql) Add(req *models.User) (usr *models.User, err error) {
	usr = new(models.User)

	query := `INSERT INTO users(first_name, last_name, email, phone_number, password) VALUES ($1, $2, $3, $4, $5) RETURNING id, first_name, last_name, email, phone_number, password, date_created, date_updated`

	err = m.conn.QueryRow(query, req.FirstName, req.LastName, req.Email, req.PhoneNumber, req.Password).Scan(&usr.Id, &usr.FirstName, &usr.LastName, &usr.Email, &usr.PhoneNumber, &usr.Password, &usr.DateCreated, &usr.DateUpdated)
	if err != nil {
		return
	}

	return
}

func (m *userSql) GetByEmail(email string) (usr *models.User, err error) {
	usr = new(models.User)

	query := `SELECT id, email, password, first_name, last_name, phone_number, date_created, date_updated FROM users WHERE email = $1`

	err = m.conn.QueryRow(query, email).Scan(&usr.Id, &usr.Email, &usr.Password, &usr.FirstName, &usr.LastName, &usr.PhoneNumber, &usr.DateCreated, &usr.DateUpdated)
	if err != nil {
		return
	}

	return
}

func (m *userSql) GetById(userId string) (usr *models.User, err error) {
	usr = new(models.User)

	query := `SELECT id, email, password, first_name, last_name, phone_number, date_created, date_updated FROM users WHERE email = $1`

	err = m.conn.QueryRow(query, userId).Scan(&usr.Id, &usr.Email, &usr.Password, &usr.FirstName, &usr.LastName, &usr.PhoneNumber, &usr.DateCreated, &usr.DateUpdated)
	if err != nil {
		return
	}

	return
}

func NewUserRepo(conn *sql.DB) UserRepo {
	return &userSql{conn: conn}
}
