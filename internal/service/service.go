package service

import (
	"context"
	"gitlab.com/ast3am77/test-go/internal/models"
)

type db interface {
	CreateOrganization(ctx context.Context, organization *models.Organization) error
	EditOrganization(ctx context.Context, organization *models.Organization) error
	GetEmployeeInfoByUserID(ctx context.Context, userId int) (*models.Employee, error)
	GetInfoOrganization(ctx context.Context, orgId int) (*models.Organization, error)
	CreateFilial(ctx context.Context, filial *models.FilialDTO) error
	EditFilial(ctx context.Context, filial *models.FilialDTO) error
	GetInfoFilial(ctx context.Context, filialId int) (*models.Filial, error)
	GetOrganizationFilials(ctx context.Context, orgId int) ([]*models.OrganizationFilialsDTO, error)
	AddWorkerFilial(ctx context.Context, token string, employee *models.AddWorkersDTO) (int, error)
	GetUserInviteLinkByID(ctx context.Context, userId int) (*models.EmployeeInvite, error)
	UpdateWorkerStatus(ctx context.Context, userId int) error
}

type logger interface {
	DebugMsg(msg string)
	ErrorMsg(msg string, err error)
}

type emailSender interface {
	SendVerificationEmail(worker *models.AddWorkersDTO, url string) error
}

type service struct {
	db          db
	log         logger
	emailSender emailSender
}

func NewService(db db, log logger, emailSender emailSender) *service {
	return &service{
		db,
		log,
		emailSender,
	}
}
