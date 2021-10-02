package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taisa831/go-ddd/application/usecase"
	"github.com/taisa831/go-ddd/domain/model"
	"github.com/taisa831/go-ddd/domain/repository"
	"github.com/taisa831/go-ddd/domain/service"
	"github.com/taisa831/go-ddd/interfaces/request"
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}


func (h *UserHandler) Create(c *gin.Context) {
	req := request.UserCreateRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.u.Create(req)
	if err != nil {
		var reErr *model.UserExistsError
		if errors.As(err, &reErr) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
