package db

import (
	"context"
	"database/sql"
	"gitlab.com/ast3am77/test-go/internal/models"
)

func (db *DB) GetEmployeeInfoByUserID(ctx context.Context, userId int) (*models.Employee, error) {
	var filialID sql.NullInt64
	var Email sql.NullString
	result := models.Employee{}
	queryOrder := `
	SELECT * FROM employee WHERE id = $1
	`
	err := db.dbConnect.QueryRow(ctx, queryOrder, userId).Scan(&result.ID, &result.OrganizationID,
		&filialID, &result.Position, &Email,
		&result.EmailConfirmationFlag)
	if err != nil {
		return nil, err
	}
	if filialID.Valid {
		result.FilialID = int(filialID.Int64)
	}
	if Email.Valid {
		result.Email = Email.String
	}
	return &result, nil
}

func (db *DB) GetUserInviteLinkByID(ctx context.Context, userId int) (*models.EmployeeInvite, error) {
	result := models.EmployeeInvite{}
	queryOrder := `
	SELECT * FROM employee_invite WHERE user_id = $1
	`
	err := db.dbConnect.QueryRow(ctx, queryOrder, userId).Scan(&result.ID, &result.UserId, &result.Token, &result.CreationDate)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) UpdateWorkerStatus(ctx context.Context, userId int) error {
	queryOrder := `
	UPDATE employee
	SET email_confirmation_flag = true
	WHERE id = $1 
	`
	_, err := db.dbConnect.Exec(ctx, queryOrder, userId)
	if err != nil {
		return err
	}

	queryOrder = `
	Delete from employee_invite
	where user_id = $1
	`

	_, err = db.dbConnect.Exec(ctx, queryOrder, userId)
	if err != nil {
		return err
	}

	return nil
}
