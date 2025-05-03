package http

import (
	"net/http"

	"github.com/CelticAlreadyUse/CMS/internal/helper"
	model "github.com/CelticAlreadyUse/CMS/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authUsecase model.UserAuthUsecase
}

func NewAuthHandler(authUsecase model.UserAuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
	}
}

var validate = validator.New()

func (handler *AuthHandler) RegisterRoute(r *gin.Engine) {
	g := r.Group("/v1/auth")
	g.POST("/login", handler.Login)
	g.POST("/register", handler.Register)
}
func (h *AuthHandler) Login(c *gin.Context) {
	var request model.LoginRequest
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("notblank", helper.NoWhitespaceOnly)
		v.RegisterValidation("lowercase", helper.Lowercase)
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	err := validate.Struct(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	token, err := h.authUsecase.Login(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Token: token,
	})
}
func (h *AuthHandler) Register(c *gin.Context) {
	var request *model.RegisterRequest
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("notblank", helper.NoWhitespaceOnly)
		v.RegisterValidation("lowercase", helper.Lowercase)
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid request format",
		})
		return
	}
	token, err := h.authUsecase.Register(c.Request.Context(), *request)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Token: token,
	})
}
