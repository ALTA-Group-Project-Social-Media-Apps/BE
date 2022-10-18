package delivery

import "github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/user/domain"

type LoginFormat struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case LoginFormat:
		cnv := i.(LoginFormat)
		return domain.Core{Username: cnv.Username, Password: cnv.Password}
	}
	return domain.Core{}
}
