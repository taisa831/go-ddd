package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taisa831/go-ddd/application/input"
	"github.com/taisa831/go-ddd/application/usecase"
	"github.com/taisa831/go-ddd/domain/repository"
	"github.com/taisa831/go-ddd/domain/service"
	"github.com/taisa831/go-ddd/interfaces/response"
)

type UserHandler struct {
	u *usecase.UserUsecase
}

func NewUserHandler(r repository.Repository, us service.UserService) UserHandler {
	return UserHandler{
		u: usecase.NewUserUsecase(r, us),
	}
}

func (h *UserHandler) List(c *gin.Context) {
	users, err := h.u.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	res := response.NewUserListResponse(users)
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Create(c *gin.Context) {
	in := input.UserCreateInput{}
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	err := h.u.Create(in)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
