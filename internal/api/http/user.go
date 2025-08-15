package server

import (
	"my-pet-simple-messenger/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAPI struct {
	service *service.UserService
}

func NewUserAPI(s *service.UserService) *UserAPI {
	return &UserAPI{service: s}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone_number" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (api *UserAPI) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := api.service.Register(req.Username, req.Surname, req.Email, req.Phone, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully"})
}
