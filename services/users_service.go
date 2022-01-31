package services

import (
	"github.com/LordRadamanthys/bookstore_users-api/domain/users"
	"github.com/LordRadamanthys/bookstore_users-api/utils/crypto_utils"
	"github.com/LordRadamanthys/bookstore_users-api/utils/date"
	"github.com/LordRadamanthys/bookstore_utils-go/rest_errors"
)

var (
	UserService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	CreateUser(user users.User) (*users.User, *rest_errors.RestErr)
	UpdateUser(isPartial bool, user users.User) (*users.User, *rest_errors.RestErr)
	GetUser(userID int) (*users.User, *rest_errors.RestErr)
	DeleteUser(userID int) *rest_errors.RestErr
	FindByStatus(status string) (users.Users, *rest_errors.RestErr)
	LoginUser(request users.LoginRequest) (*users.User, *rest_errors.RestErr)
}

func (s *usersService) CreateUser(user users.User) (*users.User, *rest_errors.RestErr) {
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

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *rest_errors.RestErr) {
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

func (s *usersService) GetUser(userID int) (*users.User, *rest_errors.RestErr) {

	result := users.User{Id: userID}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *usersService) DeleteUser(userID int) *rest_errors.RestErr {
	user, err := s.GetUser(userID)

	if err != nil {
		return err
	}
	if deleteErr := user.Delete(); deleteErr != nil {
		return deleteErr
	}

	return nil
}

func (s *usersService) FindByStatus(status string) (users.Users, *rest_errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *usersService) LoginUser(login users.LoginRequest) (*users.User, *rest_errors.RestErr) {
	dao := &users.User{
		Email:    login.Email,
		Password: crypto_utils.GetMd5(login.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}
