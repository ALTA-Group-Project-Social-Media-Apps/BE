package delivery

import (
	"net/http"

	jwt "github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/user/JWT"
	"github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/user/domain"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := userHandler{srv: srv}
  e.POST("/users", handler.AddUser())
	e.GET("/users/update", handler.updateUser())
  e.DELETE("/users", handler.DeleteByID())

}

func (us *userHandler) AddUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := us.srv.AddUser(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil register", ToResponse(res, "reg")))
	}

}

func (us *userHandler) updateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := us.srv.UpdateProfile(cnv)
    if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
    return c.JSON(http.StatusCreated, SuccessResponse("berhasil update", ToResponse(res, "reg")))
	}
  
  }
  
  func (us *userHandler) DeleteByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := jwt.ExtractToken(c)
		if id == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "cannot validate token",
			})
		}
		err := us.srv.Delete(id)
    return c.JSON(http.StatusOK, SuccessDelete("success delete user"))
	}
  }


