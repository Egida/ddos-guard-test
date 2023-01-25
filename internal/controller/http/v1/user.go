package v1

import (
	"net/http"

	"github.com/Shteyd/ddos-guard-test/internal/entity"
	"github.com/Shteyd/ddos-guard-test/internal/usecase"
	"github.com/Shteyd/ddos-guard-test/pkg/logger"
	"github.com/gin-gonic/gin"
)

type usersRoutes struct {
	useCase usecase.User
	logger  logger.Interface
}

func newUsersRoutes(handler *gin.RouterGroup, uc usecase.User, l logger.Interface) {
	r := &usersRoutes{
		useCase: uc,
		logger:  l,
	}

	handler.GET("/metrics", r.metric)
}

type metricResponse struct {
	Metric entity.Metric `json:"metric"`
}

func (r *usersRoutes) metric(ctx *gin.Context) {
	metric, err := r.useCase.Metric()
	if err != nil {
		r.logger.Error(err, "http - v1 - metric")
		errorResponse(ctx, http.StatusInternalServerError, "database problems")

		return
	}

	ctx.JSON(http.StatusOK, metricResponse{metric})
}
