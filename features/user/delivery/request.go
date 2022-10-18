package delivery

import (
	"github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/user/domain"
)

type RegisterFormat struct {
	Username string `json:"nama" form:"nama"`
	Email    string `json:"hp" form:"hp"`
	Password string `json:"password" form:"password"`
}

type UpdateFormat struct {
	ID       uint   `json:"id" form:"id"`
	Username string `json:"nama" form:"nama"`
	Email    string `json:"hp" form:"hp"`
	Photo    string `json:"photo" form:"photo"`
	Bio      string `json:"bio" form:"bio"`
	Password string `json:"password" form:"password"`
}

type LoginFormat struct {
	Username string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return domain.Core{Username: cnv.Username, Email: cnv.Email, Password: cnv.Password}
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return domain.Core{ID: cnv.ID, Username: cnv.Username, Email: cnv.Email, Photo: cnv.Photo, Bio: cnv.Bio, Password: cnv.Password}
	case LoginFormat:
		cnv := i.(LoginFormat)
		return domain.Core{Username: cnv.Username, Password: cnv.Password}
	}

	return domain.Core{}
}
