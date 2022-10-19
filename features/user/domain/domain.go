package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID       uint
	Username string
	Email    string
	Photo    string
	Bio      string
	Password string
}

type Repository interface {
	Insert(newUser Core) (Core, error)
	Update(updateData Core) (Core, error)
}

type Service interface {
	AddUser(newUser Core) (Core, error)
	UpdateProfile(updateData Core) (Core, error)
}

type Handler interface {
	AddUser() echo.HandlerFunc
	updateUser() echo.HandlerFunc
}
