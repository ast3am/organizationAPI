package db

import (
	"context"
	"database/sql"
	"fmt"
	"gitlab.com/ast3am77/test-go/internal/models"
	"time"
)

func (db *DB) CreateFilial(ctx context.Context, filial *models.FilialDTO) error {
	if filial.DirectorID > 0 {
		queryOrder := `
	INSERT INTO filial (filial_name, country, city, address, filial_type, phone_list, email_list, photo_id_list, organization_id, director_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
		_, err := db.dbConnect.Exec(ctx, queryOrder, filial.FilialName, filial.Country, filial.City, filial.Address, filial.FilialType,
			filial.PhoneList, filial.EmailList, filial.PhotoIDList, filial.OrganizationID, filial.DirectorID)
		if err != nil {
			return err
		}

		queryOrder = `
	UPDATE employee
	SET filial_id = (SELECT filial_id FROM filial WHERE director_id = $1)
	WHERE id = $1`
		_, err = db.dbConnect.Exec(ctx, queryOrder, filial.DirectorID)
		if err != nil {
			return err
		}
	} else {
		queryOrder := `
		INSERT INTO filial (filial_name, country, city, address, filial_type, phone_list, email_list, photo_id_list, organization_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
		_, err := db.dbConnect.Exec(ctx, queryOrder, filial.FilialName, filial.Country, filial.City, filial.Address, filial.FilialType,
			filial.PhoneList, filial.EmailList, filial.PhotoIDList, filial.OrganizationID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *DB) EditFilial(ctx context.Context, filial *models.FilialDTO) error {
	queryOrder := `
	UPDATE filial
	SET filial_name = $1, country = $2, city = $3, address = $4, filial_type = $5, phone_list = $6, 
	    email_list = $7, photo_id_list = $8 
	WHERE filial_id = $9 AND organization_id = (select organization_id from employee where id = $10)`
	row, err := db.dbConnect.Exec(ctx, queryOrder, filial.FilialName, filial.Country, filial.City, filial.Address,
		filial.FilialType, filial.PhoneList, filial.EmailList, filial.PhotoIDList, filial.FilialID, filial.UserID)
	if err != nil {
		return err
	}
	if row.RowsAffected() == 0 {
		err = fmt.Errorf("filial with owner_id %d not found or you not owner this filial ", filial.FilialID)
		return err
	}
	return nil
}

func (db *DB) GetInfoFilial(ctx context.Context, filialId int) (*models.Filial, error) {
	result := models.Filial{}
	var DirectorID sql.NullInt64
	queryOrder := `
	select * from filial where filial_id = $1`
	err := db.dbConnect.QueryRow(ctx, queryOrder, filialId).Scan(&result.FilialID, &result.FilialName,
		&result.Country, &result.City, &result.Address, &result.FilialType, &result.PhoneList,
		&result.EmailList, &result.PhotoIDList, &result.OrganizationID, &DirectorID)
	if err != nil {
		return nil, err
	}
	if DirectorID.Valid {
		result.DirectorID = int(DirectorID.Int64)
	}
	return &result, nil
}

func (db *DB) AddWorkerFilial(ctx context.Context, token string, employee *models.AddWorkersDTO) (int, error) {
	var filialID sql.NullInt64
	if employee.FilialID != 0 {
		filialID = sql.NullInt64{
			Int64: int64(employee.FilialID),
			Valid: true,
		}
	}
	queryOrder := `
	insert into employee (organization_id, filial_id, position, email)
	values ($1, $2, $3, $4)
	returning id`
	var newWorkerId int
	err := db.dbConnect.QueryRow(ctx, queryOrder, employee.OrganizationID, filialID, employee.Position, employee.Email).Scan(&newWorkerId)
	if err != nil {
		return 0, err
	}

	creationDate := time.Now()

	queryOrder = `
	INSERT into employee_invite (user_id, token, creation_date)
	values ($1, $2, $3)
	returning id`
	_, err = db.dbConnect.Exec(ctx, queryOrder, newWorkerId, token, creationDate)
	if err != nil {
		return 0, err
	}

	return newWorkerId, nil
}
