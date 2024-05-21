package models

import (
	"api-devbook/src/security"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/pkg/errors"
)

type User struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Alias     string    `json:"alias,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("Name is required")
	}

	if user.Alias == "" {
		return errors.New("Alias is required")
	}

	if user.Email == "" {
		return errors.New("Email is required")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("Invalid email")
	}
	

	if step == "create" &&  user.Password == "" {
		return errors.New("Password is required")
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Alias = strings.TrimSpace(user.Alias)
	user.Email = strings.TrimSpace(user.Email)

	if step == "create" {
		hashedPassword, err := security.Hash(user.Password)
		if err != nil {
			log.Println("Error hashing password: ", err)
			return err
		}
		user.Password = string(hashedPassword)
	}
	return nil
}

func (user *User) Prepare(step string) error {
	user.format(step)
	if err := user.validate(step); err != nil {
		return err
	}


	err := user.validate(step)
	if err != nil {
		return err
	}

	user.CreatedAt = time.Now()

	return nil
}