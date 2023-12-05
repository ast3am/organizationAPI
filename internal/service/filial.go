package service

import (
	"context"
	"fmt"
	"gitlab.com/ast3am77/test-go/internal/models"
	"gitlab.com/ast3am77/test-go/pkg/utils"
	"strconv"
)

func (s *service) MakeNewFilial(ctx context.Context, filial *models.FilialDTO) error {
	employee, err := s.db.GetEmployeeInfoByUserID(ctx, filial.UserID)
	if err != nil {
		return err
	}
	if employee.Position != "owner" && employee.Position != "director" || employee.OrganizationID != filial.OrganizationID || !employee.EmailConfirmationFlag {
		err = fmt.Errorf("you cannot create a filial")
		return err
	}
	if employee.Position == "director" {
		filial.DirectorID = filial.UserID
	}
	err = s.db.CreateFilial(ctx, filial)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) EditFilial(ctx context.Context, filial *models.FilialDTO) error {
	employee, err := s.db.GetEmployeeInfoByUserID(ctx, filial.UserID)

	if err != nil {
		return err
	}

	if employee.Position != "owner" && employee.Position != "director" || employee.Position == "director" && employee.FilialID != filial.FilialID || !employee.EmailConfirmationFlag {
		err = fmt.Errorf("you cannot create a filial")
		return err
	}
	err = s.db.EditFilial(ctx, filial)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetInfoFilial(ctx context.Context, filialId, userId string) (*models.Filial, error) {
	filialIdInt, err := strconv.Atoi(filialId)
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

	if employee.Position != "owner" && employee.FilialID != filialIdInt || !employee.EmailConfirmationFlag {
		err = fmt.Errorf("you cannot get info about filial with id %d", filialIdInt)
		return nil, err
	}

	result, err := s.db.GetInfoFilial(ctx, filialIdInt)
	if err != nil {
		return nil, err
	}

	if result.OrganizationID != employee.OrganizationID {
		err = fmt.Errorf("you cannot get info about filial with id %d", filialIdInt)
		return nil, err
	}

	return result, nil
}

func (s *service) AddWorkerFilial(ctx context.Context, employee *models.AddWorkersDTO, host string) error {
	checkEmployee, err := s.db.GetEmployeeInfoByUserID(ctx, employee.UserID)
	if err != nil {
		return err
	}
	if checkEmployee.Position != "owner" && checkEmployee.Position != "director" || checkEmployee.OrganizationID != employee.OrganizationID || !checkEmployee.EmailConfirmationFlag || checkEmployee.Position == "director" && checkEmployee.FilialID != employee.FilialID {
		err = fmt.Errorf("you cannot make a new worker")
		return err
	}

	token, err := utils.TokenGenerate()
	if err != nil {
		err = fmt.Errorf("cannot make a new token")
		return err
	}

	userID, err := s.db.AddWorkerFilial(ctx, token, employee)
	if err != nil {
		return err
	}

	inviteLink := fmt.Sprintf("http://%s/confirm_worker/%d/%s", host, userID, token)

	err = s.emailSender.SendVerificationEmail(employee, inviteLink)
	if err != nil {
		s.log.DebugMsg("error to send email")
		return err
	}

	return nil
}
