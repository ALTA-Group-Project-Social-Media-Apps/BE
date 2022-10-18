package delivery

import (
	"net/http"

	"github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/jwt"
	"github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/user/domain"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := userHandler{srv: srv}
	e.POST("/users/login", handler.LoginUser())
}

func (us *userHandler) LoginUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var resQry LoginFormat
		if err := c.Bind(&resQry); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}

		cnv := ToDomain(resQry)
		res, err := us.srv.Login(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		token := jwt.GenerateToken(res.ID)
		return c.JSON(http.StatusCreated, SuccessLogin("berhasil register", token, ToResponse(res, "reg")))
	}
}
