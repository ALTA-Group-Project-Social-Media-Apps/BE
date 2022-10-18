package repository

import (
	"github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/user/domain"

	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.Repository {
	return &repoQuery{
		db: dbConn,
	}
}

func (rq *repoQuery) Login(loginUser domain.Core) (domain.Core, error) {
	var resQry User
	if err := rq.db.First(&resQry, "username = ? && password", loginUser.Username, loginUser.Password).Error; err != nil {
		return domain.Core{}, err
	}

	res := ToDomain(resQry)
	return res, nil
}
