package service

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func (s *service) UpdateConfirmStatus(ctx context.Context, userId, token string) error {
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return err
	}

	employee, err := s.db.GetUserInviteLinkByID(ctx, userIdInt)
	if err != nil {
		return err
	}

	if employee.Token != token {
		err = fmt.Errorf("token failed")
		return err
	}

	ttl := 24 * time.Hour
	if time.Now().After(employee.CreationDate.Add(ttl)) {
		err = fmt.Errorf("token time off")
		return err
	}

	err = s.db.UpdateWorkerStatus(ctx, userIdInt)
	if err != nil {
		return err
	}

	return nil
}
