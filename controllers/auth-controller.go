package controllers

import (
	"auth/models"
	"auth/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthControllerInterface interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type AuthController struct {
	service *services.AuthService
}

func NewAuthController(service *services.AuthService) *AuthController {
	return &AuthController{service: service}
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var loginRequest models.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.service.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	c.JSON(http.StatusOK, models.AuthResponse{Token: token})
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var registerRequest models.RegisterRequest
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.service.Register(registerRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado exitosamente"})
}
