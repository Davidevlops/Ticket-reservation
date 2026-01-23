package repository

import (
	"Ticket-reservation-system/internal/model"
	"Ticket-reservation-system/internal/connection/database_connection"
	"errors"
)

type UserSignUpRepository struct {
	DB database_connection.DBConnection
}

func (repo *UserSignUpRepository) IsUsernameTaken(username string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM usernames WHERE Value()=$1 AND Status()='1'"
	err := repo.DB.QueryRow(query, username).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (repo *UserSignUpRepository) IsEmailTaken(email string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM emails WHERE Value()=$1 AND Status()='1'"
	err := repo.DB.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}	
func (repo *UserSignUpRepository) IsPhoneNumberTaken(phoneNumber string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM phone_numbers WHERE Value()=$1 AND Status()='1'"
	err := repo.DB.QueryRow(query, phoneNumber).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
func (repo *UserSignUpRepository) CreateUser(username model.Username, email model.Email, phoneNumber model.PhoneNumber) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var userID int
	userInsertQuery := "INSERT INTO users (Status(), created_at, updated_at) VALUES ('1', NOW(), NOW()) RETURNING id"
	err = tx.QueryRow(userInsertQuery).Scan(&userID)
	if err != nil {
		return
	}

	usernameInsertQuery := "INSERT INTO usernames (user_id, Value(), Status(), created_at, updated_at) VALUES ($1, $2, '1', NOW(), NOW())"
	_, err = tx.Exec(usernameInsertQuery, userID, username.Value())
	if err != nil {
		return err
	}

	emailInsertQuery := "INSERT INTO emails (user_id, Value(), Status(), created_at, updated_at) VALUES ($1, $2, '1', NOW(), NOW())"
	_, err = tx.Exec(emailInsertQuery, userID, email.Value())
	if err != nil {
		return err
	}

	phoneNumberInsertQuery := "INSERT INTO phone_numbers (user_id, Value(), Status(), created_at, updated_at) VALUES ($1, $2, '1', NOW(), NOW())"
	_, err = tx.Exec(phoneNumberInsertQuery, userID, phoneNumber.Value())

	if err != nil {
		return err
	}
	return nil
}
func (repo *UserSignUpRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	var userStatus, usernameStatus, emailStatus, phoneNumberStatus string
	query := `
		SELECT u.id, u.Status(), un.id, un.Value(), un.Status(), un.created_at, un.updated_at,
			   e.id, e.Value(), e.Status(), e.created_at, e.updated_at, pn.id, pn.Value(), pn.Status(), pn.created_at, pn.updated_at			
		FROM users u
		JOIN usernames un ON u.id = un.user_id
		JOIN emails e ON u.id = e.user_id		
		JOIN phone_numbers pn ON u.id = pn.user_id
		WHERE un.Value() = $1 AND u.Status() = '1'
	`
	row := repo.DB.QueryRow(query, username)	
	err := row.Scan(
		&user.ID, &userStatus,
		&user.Username.ID, &user.Username.Value(), &usernameStatus, &user.Username.CreatedAt, &user.Username.UpdatedAt,	
		&user.Email.ID, &user.Email.Value(), &emailStatus, &user.Email.CreatedAt, &user.Email.UpdatedAt,
		&user.PhoneNumber.ID, &user.PhoneNumber.Value(), &phoneNumberStatus, &user.PhoneNumber.CreatedAt, &user.PhoneNumber.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (repo *UserSignUpRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	var userStatus, usernameStatus, emailStatus, phoneNumberStatus string
	query := `
		SELECT u.id, u.Status(), un.id, un.Value(), un.Status(), un.created_at, un.updated_at,
			   e.id, e.Value(), e.Status(), e.created_at, e.updated_at, pn.id, pn.Value(), pn.Status(), pn.created_at, pn.updated_at			
		FROM users u			
		JOIN usernames un ON u.id = un.user_id
		JOIN emails e ON u.id = e.user_id		
		JOIN phone_numbers pn ON u.id = pn.user_id		
		WHERE e.Value() = $1 AND u.Status() = '1'
	`
	row := repo.DB.QueryRow(query, email)
	err := row.Scan(
		&user.ID, &userStatus,
		&user.Username.ID, &user.Username.Value(), &usernameStatus, &user.Username.CreatedAt, &user.Username.UpdatedAt,	
		&user.Email.ID, &user.Email.Value(), &emailStatus, &user.Email.CreatedAt, &user.Email.UpdatedAt,		

		&user.PhoneNumber.ID, &user.PhoneNumber.Value(), &phoneNumberStatus, &user.PhoneNumber.CreatedAt, &user.PhoneNumber.UpdatedAt,

	)
	if err != nil {
		return nil, err
	}	
	return &user, nil
}
func (repo *UserSignUpRepository) GetUserByPhoneNumber(phoneNumber string) (*model.User, error) {
	var user model.User
	var userStatus, usernameStatus, emailStatus, phoneNumberStatus string
	query := `
		SELECT u.id, u.Status(), un.id, un.Value(), un.Status(), un.created_at, un.updated_at,
			   e.id, e.Value(), e.Status(), e.created_at, e.updated_at, pn.id, pn.Value(), pn.Status(), pn.created_at, pn.updated_at			
		FROM users u			
		JOIN usernames un ON u.id = un.user_id
		JOIN emails e ON u.id = e.user_id
		JOIN phone_numbers pn ON u.id = pn.user_id
		WHERE pn.Value() = $1 AND u.Status() = '1'
	`
	row := repo.DB.QueryRow(query, phoneNumber)	
	err := row.Scan(
		&user.ID, &userStatus,
		&user.Username.ID, &user.Username.Value(), &usernameStatus, &user.Username.CreatedAt, &user.Username.UpdatedAt,	
		&user.Email.ID, &user.Email.Value(), &emailStatus, &user.Email.CreatedAt, &user.Email.UpdatedAt,		
		&user.PhoneNumber.ID, &user.PhoneNumber.Value(), &phoneNumberStatus, &user.PhoneNumber.CreatedAt, &user.PhoneNumber.UpdatedAt,
	)		
	if err != nil {
		return nil, err
	}
	return &user, nil
}