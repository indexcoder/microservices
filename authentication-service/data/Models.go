package data

import (
	"context"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models{
		User: User{},
	}
}

type Models struct {
	User User
}

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Password  string    `json:"-"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) GetAll() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, first_name, last_name, active, created_at, updated_at from users order by last_name`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Active, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (u *User) GetByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, first_name, last_name, password, active, created_at, updated_at from users where email = ?`

	var user User
	row := db.QueryRowContext(ctx, query, email)

	err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.Active, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) GetByID(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, first_name, last_name, password, active, created_at, updated_at from users where id = ?`

	var user User
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.Active, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `update users set email = $1, first_name = $2, last_name = $3, active = $4, updated_at = $5 where id = $6`

	_, err := db.ExecContext(ctx, query, u.Email, u.FirstName, u.LastName, u.Active, u.UpdatedAt, u.ID)

	return err
}

func (u *User) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `delete from users where id = $1`
	_, err := db.ExecContext(ctx, query, u.ID)
	return err
}

func (u *User) DeleteById(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `delete from users where id = $1`
	_, err := db.ExecContext(ctx, query, id)
	return err
}

func (u *User) Insert(user User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return 0, err
	}
	var newId int
	query := `insert into users (email, first_nam, last_name, password, active, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7) returning id`

	err = db.QueryRowContext(ctx, query, user.Email, user.FirstName, user.LastName, hashedPassword, user.Active, user.CreatedAt, time.Now(), time.Now()).Scan(&newId)
	if err != nil {
		return 0, err
	}

	return newId, nil
}

func (u *User) ResetPassword(password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	query := `update users set password = $1 where id = $2`
	_, err = db.ExecContext(ctx, query, hashedPassword, u.ID)
	return err
}

func (u *User) PasswordMatches(plainText string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	if err != nil {
		// Если ошибка совпадает с bcrypt.ErrMismatchedHashAndPassword
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil // Пароль неверный, но это не критическая ошибка
		}
		// Любая другая ошибка
		return false, err
	}
	return true, nil
}
