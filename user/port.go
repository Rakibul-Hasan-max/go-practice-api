package user

import (
	"go-practice-api/domain"
	userHandler "go-practice-api/rest/handlers/user"
)

type Service interface {
	userHandler.Service // Embedded interface
}

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
	// List() ([]*User, error)
	// Delete(userID int) error
	// Update(user User) (*User, error)
}
