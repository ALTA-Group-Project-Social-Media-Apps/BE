package repository

import (
	"github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/user/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Photo    string
	Bio      string
	Password string
}

func FromDomain(du domain.Core) User {
	return User{
		Model:    gorm.Model{ID: du.ID},
		Username: du.Username,
		Email:    du.Email,
		Photo:    du.Photo,
		Bio:      du.Bio,
		Password: du.Password,
	}
}

func ToDomain(u User) domain.Core {
	return domain.Core{

		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Photo:    u.Photo,
		Bio:      u.Bio,
		Password: u.Password,
	}
}

func ToDomainArray(au []User) []domain.Core {
	var res []domain.Core
	for _, val := range au {
		res = append(res, domain.Core{ID: val.ID, Username: val.Username, Email: val.Email, Photo: val.Photo, Bio: val.Bio, Password: val.Password})
	}

	return res
}
