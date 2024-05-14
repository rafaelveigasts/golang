package repository

import (
	"api-devbook/src/models"
	"database/sql"
)

type usuarios struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *usuarios {
	return &usuarios{db}
}

func (u usuarios) CreateUser(user models.User) (int64, error) {
	stmt, err := u.db.Prepare("insert into users (name, alias, email, password) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Alias, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (u usuarios) GetUsers(nameOrAlias string) ([]models.User, error) {

	nameOrAlias = "%" + nameOrAlias + "%"

	lines, err := u.db.Query("select id, name, alias, email, created_at from users where name like ? or alias like ?", nameOrAlias, nameOrAlias)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(&user.ID, &user.Name, &user.Alias, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u usuarios) GetUserByID(id int64) (
	models.User, error) {

	lines, err := u.db.Query("select id, name, alias, email, created_at from users where id = ?", id)
	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(&user.ID, &user.Name, &user.Alias, &user.Email, &user.CreatedAt); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (u usuarios) UpdateUser(id int64, user models.User) error {
	stmt, err := u.db.Prepare("update users set name = ?, alias = ?, email = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Name, user.Alias, user.Email, id); err != nil {
		return err
	}

	return nil
}

func (u usuarios) DeleteUser(id int64) error {
	stmt, err := u.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func (u usuarios) GetUserByEmail(email string) (models.User, error) {
	lines, err := u.db.Query("select id, password from users where email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}