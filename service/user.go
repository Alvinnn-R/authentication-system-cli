package service

import (
	"authentication-system-cli/model"
	"authentication-system-cli/repository"
	"authentication-system-cli/utils"
	"regexp"
	"strings"
	"unicode"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(fullName, email, phoneNumber, password string) error {
	if err := s.validateEmail(email); err != nil {
		return err
	}

	existingUser, err := s.repo.FindByEmail(email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return utils.ErrEmailExists
	}

	if err := s.validatePhoneNumber(phoneNumber); err != nil {
		return err
	}

	if err := s.validatePassword(password); err != nil {
		return err
	}

	users, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	newUser := model.User{
		FullName:    fullName,
		Email:       email,
		PhoneNumber: phoneNumber,
		Password:    password,
	}

	users = append(users, newUser)
	return s.repo.SaveAll(users)
}

func (s *UserService) Login(email, password string) (*model.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, utils.ErrUserNotFound
	}

	if user.Password != password {
		return nil, utils.ErrPasswordWrong
	}

	return user, nil
}

func (s *UserService) validateEmail(email string) error {
	email = strings.TrimSpace(email)
	if email == "" {
		return utils.ErrEmailInvalid
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(emailRegex, email)
	if !matched {
		return utils.ErrEmailInvalid
	}

	return nil
}

func (s *UserService) validatePhoneNumber(phone string) error {
	phone = strings.TrimSpace(phone)

	if len(phone) < 10 || len(phone) > 15 {
		return utils.ErrPhoneInvalid
	}

	for _, ch := range phone {
		if !unicode.IsDigit(ch) {
			return utils.ErrPhoneInvalid
		}
	}

	return nil
}

func (s *UserService) validatePassword(password string) error {
	if len(password) < 6 {
		return utils.ErrPasswordInvalid
	}
	return nil
}
