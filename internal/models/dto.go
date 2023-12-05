package models

type FilialDTO struct {
	UserID         int    `json:"user_id"`
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

type OrganizationFilialsDTO struct {
	FilialID   int    `json:"filial_id"`
	FilialName string `json:"filial_name"`
}

type AddWorkersDTO struct {
	ID                    int    `json:"id"`
	OrganizationID        int    `json:"organization_id"`
	FilialID              int    `json:"filial_id"`
	UserID                int    `json:"user_id"`
	Position              string `json:"position"`
	Email                 string `json:"email"`
	EmailConfirmationFlag bool   `json:"email_confirmation_flag"`
}

type ConfirmEmployeeDTO struct {
}
