package services

import (
	"errors"
	"strings"

	"github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/user/domain"

	"github.com/labstack/gommon/log"
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

/*
func (us *userService) ShowAllUser() ([]domain.Core, error) {
	res, err := us.qry.GetAll()
	if err != nil {
		log.Error(err.Error())
		return nil, errors.New("database error")
	}

	if len(res) == 0 {
		log.Info("no data")
		return nil, errors.New("no data")
	}

	return res, nil
}
*/

func (us *userService) Login(loginUser domain.Core) (domain.Core, error) {
	res, err := us.qry.Login(loginUser)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return domain.Core{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.Core{}, errors.New("no data")
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(loginUser.Password))
	if err != nil {
		return domain.Core{}, errors.New("password doesn't match")
	}
	return res, nil

}
