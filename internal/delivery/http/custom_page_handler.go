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

type CustomPageHandler struct {
	usecase model.CustomPageUsecase
}

func NewCustomPageHandler(usecase model.CustomPageUsecase) *CustomPageHandler {
	return &CustomPageHandler{usecase: usecase}
}

func (h *CustomPageHandler) RegisterRoute(r *gin.Engine) {
	g := r.Group("/v1/pages")
	g.GET("", h.ListPages)
	g.GET(":id", h.GetPageByID)
	g.GET("/url/:url", h.GetPageByURL)
	g.POST("", middleware.AuthMiddleware(), h.CreatePage)
	g.PUT(":id", middleware.AuthMiddleware(), h.UpdatePage)
	g.DELETE(":id", middleware.AuthMiddleware(), h.DeletePage)
}
func (h *CustomPageHandler) ListPages(c *gin.Context) {
	pages, err := h.usecase.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Data: pages, Message: "Successfully fetched custom pages"})
}

func (h *CustomPageHandler) GetPageByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid page ID"})
		return
	}
	page, err := h.usecase.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Page not found"})
		return
	}
	c.JSON(http.StatusOK, Response{Data: page, Message: "Successfully fetched page detail"})
}

func (h *CustomPageHandler) GetPageByURL(c *gin.Context) {
	url := c.Param("url")
	page, err := h.usecase.GetByURL(c.Request.Context(), url)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Page not found"})
		return
	}
	c.JSON(http.StatusOK, Response{Data: page, Message: "Successfully fetched page detail"})
}

func (h *CustomPageHandler) CreatePage(c *gin.Context) {
	var req model.CustomPageRequest
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("notblank", helper.NoWhitespaceOnly)
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request format"})
		return
	}
	claims := c.Request.Context().Value(model.BearerAuthKey).(model.CustomClaims)
	page, err := h.usecase.Create(c.Request.Context(), claims.UserID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, Response{Data: page.ID, Message: "Custom page created successfully"})
}

func (h *CustomPageHandler) UpdatePage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid page ID"})
		return
	}
	var req *model.CustomPageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request format"})
		return
	}
	claims := c.Request.Context().Value(model.BearerAuthKey).(model.CustomClaims)
	if err := h.usecase.Update(c.Request.Context(), int64(id), claims.UserID,req); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Message: "Custom page updated successfully"})
}

func (h *CustomPageHandler) DeletePage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid page ID"})
		return
	}
	if err := h.usecase.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Message: "Custom page deleted successfully"})
}
