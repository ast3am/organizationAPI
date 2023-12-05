package service

import (
	"context"
	"fmt"
	"gitlab.com/ast3am77/test-go/internal/models"
	"strconv"
)

func (s *service) MakeNewOrganization(ctx context.Context, organization *models.Organization) error {
	err := s.db.CreateOrganization(ctx, organization)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) EditOrganization(ctx context.Context, organization *models.Organization) error {
	err := s.db.EditOrganization(ctx, organization)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetInfoOrganization(ctx context.Context, orgId, userId string) (*models.Organization, error) {
	orgInt, err := strconv.Atoi(orgId)
	if err != nil {
		return nil, err
	}
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return nil, err
	}
	employee, err := s.db.GetEmployeeInfoByUserID(ctx, userIdInt)
	if err != nil {
		return nil, err
	}
	if employee.OrganizationID != orgInt {
		err = fmt.Errorf("access denied because user with id: %d don't belong to organization with id: %d", userIdInt, orgInt)
		return nil, err
	}

	result := &models.Organization{}
	result, err = s.db.GetInfoOrganization(ctx, orgInt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *service) GetOrganizationsFilials(ctx context.Context, orgId, userId string) ([]*models.OrganizationFilialsDTO, error) {
	orgInt, err := strconv.Atoi(orgId)
	if err != nil {
		return nil, err
	}
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return nil, err
	}

	employee, err := s.db.GetEmployeeInfoByUserID(ctx, userIdInt)
	if err != nil {
		return nil, err
	}
	if employee.OrganizationID != orgInt {
		err = fmt.Errorf("access denied because user with id: %d don't belong to organization with id: %d", userIdInt, orgInt)
		return nil, err
	}

	result, err := s.db.GetOrganizationFilials(ctx, orgInt)
	if err != nil {
		return nil, err
	}
	return result, nil
}
