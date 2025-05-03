package http

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/CelticAlreadyUse/CMS/internal/helper"
	"github.com/CelticAlreadyUse/CMS/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type CommentHandler struct {
	commentUsecase model.CommentUsecases
}

func NewCommentsHandler(usecase model.CommentUsecases) *CommentHandler {
	return &CommentHandler{commentUsecase: usecase}
}
func (handler *CommentHandler) RegisterRoute(r *gin.Engine) {
	g := r.Group("/v1/comments")
	g.POST("", handler.CreateComment)
	g.DELETE("", handler.DeleteComment)
}
func (h *CommentHandler) CreateComment(c *gin.Context) {
	var comment model.CommentRequest
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("notblank", helper.NoWhitespaceOnly)
	}
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid request format",
		})
		return
	}
	comment.Username = strings.TrimSpace(comment.Username)
	if comment.Username == "" {
		comment.Username =  "anonymous"
	}
	result, err := h.commentUsecase.Create(c.Request.Context(), &comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Response{
		Data:    result,
		Message: "Comment created successfully",
	})
}
func (h *CommentHandler) DeleteComment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid comment ID",
		})
		return
	}

	err = h.commentUsecase.Delete(c.Request.Context(), int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{Message: "Comment deleted successfully"})
}
