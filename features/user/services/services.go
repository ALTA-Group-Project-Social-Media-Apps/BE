package services

import (
	"errors"
	"strings"

	"github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/config"
	"github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/user/domain"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &userService{
		qry: repo,
	}
}

func (us *userService) AddUser(newUser domain.Core) (domain.Core, error) {
	var temp string
	if newUser.Password != "" {
		generate, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
		newUser.Password = string(generate)
	}
	if newUser.Password == "" {
		temp = newUser.Password
	} else {
		temp = "error"
	}
	res, err := us.qry.Insert(newUser)

	if err != nil {
		if temp == "" {
			return domain.Core{}, errors.New("cannot encript password")
		}
		return domain.Core{}, errors.New(config.DUPLICATED_DATA)
	}

	return res, nil
}
func (us *userService) UpdateProfile(updatedData domain.Core) (domain.Core, error) {
	if updatedData.Password != "" {
		generate, _ := bcrypt.GenerateFromPassword([]byte(updatedData.Password), 10)
		updatedData.Password = string(generate)
	}

	res, err := us.qry.Update(updatedData)
	if err != nil {
		if strings.Contains(err.Error(), "column") {
			return domain.Core{}, errors.New("rejected from database")
		}
		return domain.Core{}, errors.New("rejected from database")
	}

	return res, nil
}
