package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/vivaldy22/go-unit-test/6-mock/usecase"
)

type userHandler struct {
	usecase usecase.UserService
}

func NewUserHandler(usecase usecase.UserService) *userHandler {
	return &userHandler{
		usecase: usecase,
	}
}

func (u *userHandler) GetByID(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = c.JSON(http.StatusBadRequest, err)
		return
	}

	response, err := u.usecase.GetByID(int64(id))
	if err != nil {
		err = c.JSON(http.StatusInternalServerError, err)
		return
	}

	return c.JSON(http.StatusOK, response)
}
