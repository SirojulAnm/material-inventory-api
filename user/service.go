package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input InputRegister) (User, error)
	Login(input LoginAdminRequest) (User, error)
	GetUserByID(ID int) (User, error)
	WarehouseOfficer() ([]User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Register(input InputRegister) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.PasswordHash = input.Password
	user.Role = input.Role

	cekUser, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return user, err
	}

	if cekUser.ID > 0 {
		return user, errors.New("Email is already registered")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginAdminRequest) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("Email tidak ditemukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("Tidak ditemukan user berdasakan ID ini")
	}

	return user, nil
}

func (s *service) WarehouseOfficer() ([]User, error) {
	user, err := s.repository.WarehouseOfficer()
	if err != nil {
		return user, err
	}

	return user, nil
}
