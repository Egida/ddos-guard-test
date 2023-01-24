package v1

import (
	"net/http"

	"github.com/Shteyd/ddos-guard-test/internal/usecase"
	"github.com/Shteyd/ddos-guard-test/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, uc usecase.User, l logger.Interface) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.GET("/ping", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })

	handler.Group("/v1")
	{
		newUsersRoutes(handler, uc, l)
	}
}
