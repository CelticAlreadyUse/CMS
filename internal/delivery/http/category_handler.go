package http

import (
	"net/http"
	"strconv"

	"github.com/CelticAlreadyUse/CMS/internal/helper"
	"github.com/CelticAlreadyUse/CMS/internal/middleware"
	"github.com/CelticAlreadyUse/CMS/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type CategoryHandler struct {
	Categoryusecase model.CategoryUsecases
}

func InitCateoryHandler(usecase model.CategoryUsecases) *CategoryHandler {
	return &CategoryHandler{Categoryusecase: usecase}
}
func (handler *CategoryHandler) RegisterRoute(r *gin.Engine) {
	g := r.Group("/v1/categories")
	g.GET("", handler.GetCategoryList)
	g.GET("/:id", handler.GetCategoryByID)
	g.PUT("/:id", middleware.AuthMiddleware(), handler.UpdateCategory)
	g.POST("", middleware.AuthMiddleware(), handler.CreateCategory)
	g.DELETE("/:id", middleware.AuthMiddleware(), handler.DeleteCategory)

}
func (handler *CategoryHandler) GetCategoryList(c *gin.Context) {
	categories, err := handler.Categoryusecase.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, categories)
}

func (handler *CategoryHandler) GetCategoryByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "invalid id",
		})
		return
	}
	category, err := handler.Categoryusecase.GetByID(c.Request.Context(), int64(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{
		Data:    category,
		Message: "sucessfully get category",
	})
}

func (handler *CategoryHandler) CreateCategory(c *gin.Context) {
	var req model.Category
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("lowercase", helper.Lowercase)
		v.RegisterValidation("notblank", helper.NoWhitespaceOnly)
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid request format",
		})
		return
	}
	category, err := handler.Categoryusecase.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Response{
		Data:    category,
		Message: "sucessfully create category",
	})
}

func (handler *CategoryHandler) UpdateCategory(c *gin.Context) {
	idParam := c.Param("id")
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("lowercase", helper.Lowercase)
		v.RegisterValidation("notblank", helper.NoWhitespaceOnly)
	}
	var req model.Category
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "invalid request format",
		})
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "failed to convert id type",
		})
		return
	}
	category, err := handler.Categoryusecase.Update(c.Request.Context(), int64(id), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "failed to update category",
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Data:    category,
		Message: "category sucesfully updated",
	})
}

func (handler *CategoryHandler) DeleteCategory(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	if err := handler.Categoryusecase.Delete(c.Request.Context(), int64(id)); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Message: "category Sucessfully deleted",
	})
}
