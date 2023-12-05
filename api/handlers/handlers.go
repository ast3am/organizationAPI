package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"gitlab.com/ast3am77/test-go/internal/models"
	"net/http"
)

type logger interface {
	HandlerLog(r *http.Request, status int, msg string)
	HandlerErrorLog(r *http.Request, status int, msg string, err error)
	ErrorMsg(msg string, err error)
}

type services interface {
	MakeNewOrganization(ctx context.Context, organization *models.Organization) error
	EditOrganization(ctx context.Context, organization *models.Organization) error
	GetInfoOrganization(ctx context.Context, orgId, userId string) (*models.Organization, error)
	MakeNewFilial(ctx context.Context, filial *models.FilialDTO) error
	EditFilial(ctx context.Context, filial *models.FilialDTO) error
	GetInfoFilial(ctx context.Context, filialId, userId string) (*models.Filial, error)
	GetOrganizationsFilials(ctx context.Context, orgId, userId string) ([]*models.OrganizationFilialsDTO, error)
	AddWorkerFilial(ctx context.Context, employee *models.AddWorkersDTO, host string) error
	UpdateConfirmStatus(ctx context.Context, userId, token string) error
}

type Handler struct {
	services services
	log      logger
}

func NewHandler(serv services, log logger) *Handler {
	return &Handler{
		services: serv,
		log:      log,
	}
}

func (h *Handler) RegisterHandlers(router *gin.Engine) {
	organization := router.Group("/v1/organization")
	{
		organization.POST("/create", h.CreateOrganization)
		organization.POST("/edit", h.EditOrganization)
		organization.GET("/get_info/:org_id/:user_id", h.GetInfoOrganization)
		organization.GET("filials/:org_id/:user_id", h.GetOrganizationFilials)
	}
	filial := router.Group("v1/filial")
	{
		filial.POST("/create", h.CreateFilial)
		filial.POST("/edit", h.EditFilial)
		filial.GET("/get_info/:filial_id/:user_id", h.GetInfoFilial)
		filial.POST("/add_worker", h.AddWorkerFilial)
	}
	router.GET("/confirm_worker/:user_id/:token", h.UpdateConfirmStatus)
}
