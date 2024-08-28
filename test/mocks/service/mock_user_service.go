package mock_service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rentease/internal/domain/entities"
)

type MockUserService struct {
}

func NewMockUserService() *MockUserService {
	return &MockUserService{}
}

func (ms *MockUserService) SignUp(user entities.User) error {

	return nil

}

func (ms *MockUserService) FindByUsername(username string) (entities.User, error) {

	return entities.User{}, nil
}

func (ms *MockUserService) Login(username, password string) (bool, error) {

	return true, nil

}

func (ms *MockUserService) AddToWishlist(username string, propertyID primitive.ObjectID) error {
	return nil
}

func (ms *MockUserService) UpdateUser(user entities.User) error {
	return nil
}

func (ms *MockUserService) GetAllUsers() ([]entities.User, error) {
	return []entities.User{}, nil
}

func (ms *MockUserService) DeleteUser(username string) error {
	return nil
}
