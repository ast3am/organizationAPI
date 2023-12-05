package db

import (
	"context"
	"fmt"
	"gitlab.com/ast3am77/test-go/internal/models"
)

func (db *DB) CreateOrganization(ctx context.Context, organization *models.Organization) error {
	queryOrder := `
	INSERT INTO organization (organization_name, legal_type, legal_address, inn, owner_id)
	VALUES ($1, $2, $3, $4, $5)`
	_, err := db.dbConnect.Exec(ctx, queryOrder, organization.Name, organization.LegalType, organization.LegalAddress,
		organization.INN, organization.OwnerID)
	if err != nil {
		return err
	}
	queryOrder = `
	INSERT INTO employee (organization_id, id, position, email_confirmation_flag)
	VALUES (
    (SELECT org.id FROM organization org WHERE org.owner_id = $1),
    $1, 'owner', true );`
	_, err = db.dbConnect.Exec(ctx, queryOrder, organization.OwnerID)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) EditOrganization(ctx context.Context, organization *models.Organization) error {
	queryOrder := `
	UPDATE organization
	SET organization_name = $1, legal_type = $2, legal_address = $3
	WHERE owner_id = $4`
	row, err := db.dbConnect.Exec(ctx, queryOrder, organization.Name, organization.LegalType,
		organization.LegalAddress, organization.OwnerID)
	if err != nil {
		return err
	}
	if row.RowsAffected() == 0 {
		err = fmt.Errorf("organization with owner_id %d not found", organization.OwnerID)
		return err
	}
	return nil
}

func (db *DB) GetInfoOrganization(ctx context.Context, orgId int) (*models.Organization, error) {
	result := models.Organization{}
	queryOrder := `
	SELECT * FROM organization WHERE id = $1`
	err := db.dbConnect.QueryRow(ctx, queryOrder, orgId).Scan(&result.OrganizationId, &result.Name,
		&result.LegalType, &result.LegalAddress, &result.INN, &result.OwnerID)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) GetOrganizationFilials(ctx context.Context, orgId int) ([]*models.OrganizationFilialsDTO, error) {
	result := make([]*models.OrganizationFilialsDTO, 0)
	queryOrder := `
	SELECT filial_id, filial_name from filial where organization_id = $1`
	rows, err := db.dbConnect.Query(ctx, queryOrder, orgId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		temp := models.OrganizationFilialsDTO{}
		err = rows.Scan(&temp.FilialID, &temp.FilialName)
		if err != nil {
			return nil, err
		}
		result = append(result, &temp)
	}

	return result, nil
}
