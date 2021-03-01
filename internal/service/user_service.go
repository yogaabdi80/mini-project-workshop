package service

import (
	"fmt"
	"workshop-golang-psm/food-app-service/internal/domain/entity"
	"workshop-golang-psm/food-app-service/internal/domain/repository"
)

type UserService struct {
	userRepo repository.IUserRepository
}

type IUserService interface {
	SaveUser(*entity.ReqisterViewModel) (*entity.UserViewModel, error)
	GetListUser() (*[]entity.UserViewModel, error)
	GetDetailUser(id int) (*entity.UserViewModel, error)
	UpdateUser(userVM *entity.User) (*entity.User, error)
	DeleteUser(id int) error
}

func NewUserService(userRepo repository.IUserRepository) *UserService {
	var userService = UserService{}
	userService.userRepo = userRepo
	return &userService
}

func (s *UserService) GetListUser() (*[]entity.UserViewModel, error) {
	result, err := s.userRepo.GetAllUser()
	if err != nil {
		return nil, err
	}

	var users []entity.UserViewModel
	for _, item := range result {
		var user entity.UserViewModel
		user.Email = item.Email
		user.FullName = fmt.Sprintf("%s %s", item.FirstName, item.LastName)
		user.Email = item.Email
		users = append(users, user)
	}

	return &users, nil
}

func (s *UserService) GetDetailUser(id int) (*entity.UserViewModel, error) {
	var viewModel entity.UserViewModel

	result, err := s.userRepo.GetDetailUser(id)
	if err != nil {
		return nil, err
	}

	if result != nil {
		viewModel = entity.UserViewModel{
			ID:       result.ID,
			FullName: fmt.Sprintf("%s %s", result.FirstName, result.LastName),
			Email:    result.Email,
		}
	}

	return &viewModel, nil
}

func (s *UserService) SaveUser(userVM *entity.ReqisterViewModel) (*entity.UserViewModel, error) {
	var user = entity.User{
		FirstName: userVM.FirstName,
		LastName:  userVM.LastName,
		Email:     userVM.Email,
	}

	password, err := user.EncryptPassword(userVM.Password)
	if err != nil {
		return nil, err
	}

	user.Password = password

	result, err := s.userRepo.SaveUser(&user)
	if err != nil {
		return nil, err
	}

	var afterRegVM entity.UserViewModel

	if result != nil {
		afterRegVM = entity.UserViewModel{
			ID:       result.ID,
			FullName: fmt.Sprintf("%s %s", result.FirstName, result.LastName),
			Email:    result.Email,
		}
	}

	return &afterRegVM, nil
}

func (s *UserService) UpdateUser(userVM *entity.User) (*entity.User, error) {
	result, err := s.UpdateUser(userVM)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (s *UserService) DeleteUser(id int) error {
	err := s.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}
