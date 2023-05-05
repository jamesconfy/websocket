package handlers

import (
	"project-name/internal/response"
	"project-name/internal/service"

	"github.com/gin-gonic/gin"
)

type HomeHandler interface {
	Home(c *gin.Context)
}

type homeHandler struct {
	homeSrv service.HomeService
}

func (h *homeHandler) Home(c *gin.Context) {
	message, err := h.homeSrv.CreateHome()
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "", message)
}

func NewHomeHandler(homeSrv service.HomeService) HomeHandler {
	return &homeHandler{homeSrv: homeSrv}
}
