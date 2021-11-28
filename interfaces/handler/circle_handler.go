package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taisa831/go-ddd/application/usecase"
	"github.com/taisa831/go-ddd/domain/factory"
	"github.com/taisa831/go-ddd/domain/repository"
	"github.com/taisa831/go-ddd/domain/service"
)

type CircleHandler struct {
	u *usecase.CircleUsecase
}

func NewCircleHandler(r repository.Repository, f factory.CircleFactory, cs service.CircleService) CircleHandler {
	return CircleHandler{
		u: usecase.NewCircleUsecase(f, cs, r),
	}
}

func (h *CircleHandler) Create(c *gin.Context) {
	userID := c.Query("userId")
	if userID == "" {
		c.JSON(http.StatusInternalServerError, errors.New("userId is required"))
		return
	}

	circleName := c.Query("circleName")
	if circleName == "" {
		c.JSON(http.StatusInternalServerError, errors.New("circleName is required"))
		return
	}

	circle, err := h.u.Create(userID, circleName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, circle)
}

func (h *CircleHandler) Join(c *gin.Context) {
	userID := c.Query("userId")
	if userID == "" {
		c.JSON(http.StatusInternalServerError, errors.New("userId is required"))
		return
	}

	circleID := c.Param("circleId")
	if circleID == "" {
		c.JSON(http.StatusInternalServerError, errors.New("circleId is required"))
		return
	}

	err := h.u.Join(userID, circleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
