package services

import "github.com/awwabi/monorepo-academics/cuti/internal/repositories"

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{userRepo}
}

// Implement service methods here
