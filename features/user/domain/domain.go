package domain

import "github.com/labstack/echo"

type Core struct {
	ID       uint
	Username string
	Email    string
	Password string
	Photo    string
}

type Repository interface {
	Login(User Core) (Core, error)
}

type Service interface {
	Login(loginUser Core) (Core, error)
}

type Handler interface {
	Login() echo.HandlerFunc
}
