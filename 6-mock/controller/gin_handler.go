package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vivaldy22/go-unit-test/6-mock/usecase"
)

type userGinHandler struct {
	service usecase.UserService
}

type UserGinHandlerConfig struct {
	R           *gin.Engine
	UserService usecase.UserService
}

func NewUserGinHandler(conf *UserGinHandlerConfig) *userGinHandler {
	return &userGinHandler{
		service: conf.UserService,
	}
}

func (h *userGinHandler) GetByID(c *gin.Context) {
	parID, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "general error",
		})
		return
	}

	id := parID.(int64)
	response, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": response,
	})
}
