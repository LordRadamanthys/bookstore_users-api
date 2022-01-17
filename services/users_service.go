package services

import (
	"github.com/LordRadamanthys/bookstore_users-api/domain/users"
	"github.com/LordRadamanthys/bookstore_users-api/utils/crypto_utils"
	"github.com/LordRadamanthys/bookstore_users-api/utils/date"
	"github.com/LordRadamanthys/bookstore_users-api/utils/errors"
)

var (
	UserService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	CreateUser(user users.User) (*users.User, *errors.RestErr)
	UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr)
	GetUser(userID int) (*users.User, *errors.RestErr)
	DeleteUser(userID int) *errors.RestErr
	FindByStatus(status string) (users.Users, *errors.RestErr)
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = date.GetNowDBFormat()
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := s.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email

	}

	updateErr := current.Update()
	if updateErr != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService) GetUser(userID int) (*users.User, *errors.RestErr) {

	result := users.User{Id: userID}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *usersService) DeleteUser(userID int) *errors.RestErr {
	user, err := s.GetUser(userID)

	if err != nil {
		return err
	}
	if deleteErr := user.Delete(); deleteErr != nil {
		return deleteErr
	}

	return nil
}

func (s *usersService) FindByStatus(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
