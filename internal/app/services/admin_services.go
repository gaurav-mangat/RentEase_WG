package services

import "rentease/internal/domain/interfaces"

type AdminService struct {
	userRepo interfaces.UserRepo
}

func NewAdminService(userRepo interfaces.UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}
