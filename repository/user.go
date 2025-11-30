package repository

import (
	"authentication-system-cli/model"
	"authentication-system-cli/utils"
)

type UserRepository struct {
	FilePath string
}

func NewUserRepository(path string) *UserRepository {
	return &UserRepository{FilePath: path}
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	if err := utils.ReadJSON(r.FilePath, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) SaveAll(users []model.User) error {
	return utils.WriteJSON(r.FilePath, users)
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	users, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, nil
}
