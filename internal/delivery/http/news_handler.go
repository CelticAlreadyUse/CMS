package http

import (
	"net/http"

	"github.com/CelticAlreadyUse/CMS/internal/helper"
	"github.com/CelticAlreadyUse/CMS/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type NewsHandler struct {
	newsUsecase model.NewsUsecase
}

func NewNewsHandler(newsUsecase model.NewsUsecase) *NewsHandler {
	return &NewsHandler{
		newsUsecase: newsUsecase,
	}
}
func (handler *NewsHandler) RegisterRoute(r *gin.Engine) {
	g := r.Group("/v1/news")
	g.GET("", handler.GetAll)
	g.GET("/:id", handler.GetNewsByID)
	g.PUT("/:id", handler.UpdateNews)
	g.POST("", handler.CreateNews)
	g.DELETE("/:id", handler.DeleteNews)
}
func (h *NewsHandler) GetAll(c *gin.Context) {
	news, err := h.newsUsecase.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, news)
}
func (h *NewsHandler) GetNewsByID(c *gin.Context) {
	id := c.Param("id")
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("notblank", helper.NoWhitespaceOnly)
	}
	news, err := h.newsUsecase.GetByID(c.Request.Context(), id)
	if err != nil {
		helper.NotFound(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, news)
}
func (h *NewsHandler) CreateNews(c *gin.Context) {
	var req model.NewsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid request format",
		})
		return
	}

	news, err := h.newsUsecase.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"news_id": news.ID})
}
func (h *NewsHandler) UpdateNews(c *gin.Context) {
	id := c.Param("id")
	var req *model.News
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid request format",
		})
		return
	}
	err := h.newsUsecase.Update(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError,ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Message: "News Sucessfully updated",
	})
}
func (h *NewsHandler) DeleteNews(c *gin.Context) {
	id := c.Param("id")
	err := h.newsUsecase.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "news sucessfullt deleted "})
}
