package service_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"rentease/internal/app/services"
	"rentease/internal/domain/entities"
	"rentease/pkg/utils"
	mocks_interfaces "rentease/test/mocks/repository"
	"testing"
)

var (
	ctrl         *gomock.Controller
	mockUserRepo *mocks_interfaces.MockUserRepo
	userService  *services.UserService
)

func setup(t *testing.T) func() {
	// Set up the gomock controller
	ctrl = gomock.NewController(t)

	// Create a mock UserRepository
	mockUserRepo = mocks_interfaces.NewMockUserRepo(ctrl)

	// Initialize the UserService with the mock repository
	userService = services.NewUserService(mockUserRepo)

	// Return a cleanup function to be called at the end of the test
	return func() {
		ctrl.Finish()
	}
}

func TestUserService_SignUp(t *testing.T) {
	tests := []struct {
		name          string
		newUser       entities.User
		mockFindUser  *entities.User
		mockFindError error
		mockSaveError error
		expectedError bool
	}{
		{
			name: "Successful signup",
			newUser: entities.User{
				Username:     "newuser",
				PasswordHash: "hashedpassword",
				Name:         "New User",
				Age:          25,
				Email:        "newuser@example.com",
				PhoneNumber:  "1234567890",
				Address:      "123 New St",
				Role:         "tenant",
				Wishlist:     []primitive.ObjectID{},
			},
			mockFindUser:  nil, // User not found, so this should be nil
			mockFindError: mongo.ErrNoDocuments,
			mockSaveError: nil,
			expectedError: false,
		},
		{
			name: "Error during user save",
			newUser: entities.User{
				Username:     "newuser",
				PasswordHash: "hashedpassword",
				Name:         "New User",
				Age:          25,
				Email:        "newuser@example.com",
				PhoneNumber:  "1234567890",
				Address:      "123 New St",
				Role:         "tenant",
				Wishlist:     []primitive.ObjectID{},
			},
			mockFindUser:  nil, // User not found, so this should be nil
			mockFindError: mongo.ErrNoDocuments,
			mockSaveError: errors.New("save error"),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			// Mock FindByUsername behavior
			mockUserRepo.EXPECT().FindByUsername(gomock.Any(), tt.newUser.Username).Return(tt.mockFindUser, tt.mockFindError).Times(1)

			// If no existing user is found, mock SaveUser behavior and check if correct user data is passed
			if tt.mockFindError == mongo.ErrNoDocuments {
				mockUserRepo.EXPECT().SaveUser(tt.newUser).Return(tt.mockSaveError).Times(1)
			}

			err := userService.SignUp(tt.newUser)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserService_FindByUsername(t *testing.T) {
	tests := []struct {
		name          string
		username      string
		mockUser      *entities.User
		mockRepoError error
		expectedError bool
	}{
		{
			name:     "User Found",
			username: "testuser",
			mockUser: &entities.User{
				Username: "testuser",
				Email:    "testuser@example.com",
			},
			mockRepoError: nil,
			expectedError: false,
		},
		{
			name:          "User Not Found",
			username:      "nonexistentuser",
			mockUser:      nil,
			mockRepoError: nil,
			expectedError: false,
		},
		{
			name:          "Repository Error",
			username:      "testuser",
			mockUser:      nil,
			mockRepoError: errors.New("repository error"),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			mockUserRepo.EXPECT().FindByUsername(gomock.Any(), tt.username).Return(tt.mockUser, tt.mockRepoError).Times(1)

			user, err := userService.FindByUsername(tt.username)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Equal(t, entities.User{}, user)
			} else {
				assert.NoError(t, err)
				if tt.mockUser != nil {
					assert.Equal(t, *tt.mockUser, user)
				} else {
					assert.Equal(t, entities.User{}, user)
				}
			}
		})
	}
}

func TestUserService_Login(t *testing.T) {
	tests := []struct {
		name          string
		username      string
		password      string
		mockExist     bool
		mockRepoError error
		expectedError bool
		expectedExist bool
	}{
		{
			name:          "Successful Login",
			username:      "testuser",
			password:      "correctpassword",
			mockExist:     true,
			mockRepoError: nil,
			expectedError: false,
			expectedExist: true,
		},
		{
			name:          "Wrong Password",
			username:      "testuser",
			password:      "wrongpassword",
			mockExist:     false,
			mockRepoError: nil,
			expectedError: false,
			expectedExist: false,
		},
		{
			name:          "Repository Error",
			username:      "testuser",
			password:      "anyPassword",
			mockExist:     false,
			mockRepoError: errors.New("repository error"),
			expectedError: true,
			expectedExist: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			mockUserRepo.EXPECT().CheckPassword(gomock.Any(), tt.username, tt.password).Return(tt.mockExist, tt.mockRepoError).Times(1)

			success, err := userService.Login(tt.username, tt.password)

			if tt.expectedError {
				assert.Error(t, err)
				assert.False(t, success)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedExist, success)
				if success {
					assert.Equal(t, tt.username, utils.ActiveUser)
				}
			}
		})
	}
}

func TestUserService_AddToWishlist(t *testing.T) {
	tests := []struct {
		name          string
		username      string
		propertyID    primitive.ObjectID
		mockUser      *entities.User
		mockRepoError error
		expectedError bool
	}{
		{
			name:       "Successful Add to Wishlist",
			username:   "testuser",
			propertyID: primitive.NewObjectID(),
			mockUser: &entities.User{
				Username: "testuser",
				Wishlist: []primitive.ObjectID{},
			},
			mockRepoError: nil,
			expectedError: false,
		},
		{
			name:       "Property Already in Wishlist",
			username:   "testuser",
			propertyID: primitive.NewObjectID(),
			mockUser: &entities.User{
				Username: "testuser",
				Wishlist: []primitive.ObjectID{primitive.NewObjectID()},
			},
			mockRepoError: nil,
			expectedError: true,
		},
		{
			name:          "User Not Found",
			username:      "testuser",
			propertyID:    primitive.NewObjectID(),
			mockUser:      nil,
			mockRepoError: errors.New("user not found"),
			expectedError: true,
		},
		{
			name:       "Update User Error",
			username:   "testuser",
			propertyID: primitive.NewObjectID(),
			mockUser: &entities.User{
				Username: "testuser",
				Wishlist: []primitive.ObjectID{},
			},
			mockRepoError: nil,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			// Setup the expectations for FindByUsername
			mockUserRepo.EXPECT().FindByUsername(gomock.Any(), tt.username).Return(tt.mockUser, tt.mockRepoError).Times(1)

			// Setup the expectations for UpdateUser if needed
			if tt.mockUser != nil && tt.mockRepoError == nil && !tt.expectedError {
				mockUserRepo.EXPECT().UpdateUser(gomock.Any()).Return(nil).Times(1)
			} else if tt.mockRepoError == nil && tt.expectedError {
				mockUserRepo.EXPECT().UpdateUser(gomock.Any()).Return(errors.New("update user error")).Times(1)
			}

			err := userService.AddToWishlist(tt.username, tt.propertyID)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	tests := []struct {
		name          string
		user          entities.User
		mockRepoError error
		expectedError bool
	}{
		{
			name: "Successful Update",
			user: entities.User{
				Username: "testuser",
				// Other fields if necessary
			},
			mockRepoError: nil,
			expectedError: false,
		},
		{
			name: "Update User Error",
			user: entities.User{
				Username: "testuser",
				// Other fields if necessary
			},
			mockRepoError: errors.New("update user error"),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			// Set up the mock expectation
			mockUserRepo.EXPECT().UpdateUser(tt.user).Return(tt.mockRepoError).Times(1)

			// Call the UpdateUser method
			err := userService.UpdateUser(tt.user)

			// Assert the results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserService_GetAllUsers(t *testing.T) {
	tests := []struct {
		name          string
		mockUsers     []entities.User
		mockRepoError error
		expectedError bool
		expectedUsers []entities.User
	}{
		{
			name: "Successful Retrieval",
			mockUsers: []entities.User{
				{Username: "user1"},
				{Username: "user2"},
			},
			mockRepoError: nil,
			expectedError: false,
			expectedUsers: []entities.User{
				{Username: "user1"},
				{Username: "user2"},
			},
		},
		{
			name:          "Repository Error",
			mockUsers:     nil,
			mockRepoError: errors.New("repository error"),
			expectedError: true,
			expectedUsers: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			// Set up the mock expectation
			mockUserRepo.EXPECT().FindAll().Return(tt.mockUsers, tt.mockRepoError).Times(1)

			// Call the GetAllUsers method
			users, err := userService.GetAllUsers()

			// Assert the results
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, users)
			} else {
				assert.NoError(t, err)
				assert.ElementsMatch(t, tt.expectedUsers, users)
			}
		})
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	tests := []struct {
		name          string
		username      string
		mockRepoError error
		expectedError bool
	}{
		{
			name:          "Successful Deletion",
			username:      "testuser",
			mockRepoError: nil,
			expectedError: false,
		},
		{
			name:          "Repository Error",
			username:      "testuser",
			mockRepoError: errors.New("repository error"),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setup(t)
			defer teardown()

			// Set up the mock expectation
			mockUserRepo.EXPECT().Delete(tt.username).Return(tt.mockRepoError).Times(1)

			// Call the DeleteUser method
			err := userService.DeleteUser(tt.username)

			// Assert the results
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
