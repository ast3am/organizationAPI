package models

import (
	"time"
)

type Organization struct {
	OrganizationId int
	Name           string `json:"name"`
	LegalType      string `json:"legal_type"`
	LegalAddress   string `json:"legal_address"`
	INN            string `json:"inn"`
	OwnerID        int    `json:"owner_id"`
}

type Filial struct {
	FilialID       int    `json:"filial_id"`
	FilialName     string `json:"filial_name"`
	Country        string `json:"country"`
	City           string `json:"city"`
	Address        string `json:"address"`
	FilialType     string `json:"filial_type"`
	PhoneList      string `json:"phone_list"`
	EmailList      string `json:"email_list"`
	PhotoIDList    string `json:"photo_id_list"`
	OrganizationID int    `json:"organization_id"`
	DirectorID     int    `json:"director_id"`
}

type Employee struct {
	ID                    int    `json:"id"`
	OrganizationID        int    `json:"organization_id"`
	FilialID              int    `json:"filial_id"`
	Position              string `json:"position"`
	Email                 string `json:"email"`
	EmailConfirmationFlag bool   `json:"email_confirmation_flag"`
}

type EmployeeInvite struct {
	ID           int
	UserId       int
	Token        string
	CreationDate time.Time
}
