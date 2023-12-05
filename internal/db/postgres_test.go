package db

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.com/ast3am77/test-go/internal/db/mocks"
	"gitlab.com/ast3am77/test-go/internal/models"
	"testing"
)

var testCfg = models.Config{
	ListenPort: "",
	SqlConfig: models.SqlConfig{
		UsernameDB: "postgres",
		PasswordDB: "password",
		HostDB:     "localhost",
		PortDB:     "5430",
		DBName:     "organization_db",
		DelayTime:  10,
	},
	EmailDealerConfig: models.EmailDealerConfig{Host: "", Port: 0, Username: "", Password: ""},
	LogLevel:          "",
}

func TestNewClient(t *testing.T) {
	ctx := context.Background()
	log := mocks.NewLogger(t)

	testTable := []struct {
		name string
		cfg  models.Config
	}{
		{
			name: "positive",
			cfg:  testCfg,
		},
		{
			name: "negative",
			cfg: models.Config{
				ListenPort: "",
				SqlConfig: models.SqlConfig{
					UsernameDB: "postgres",
					PasswordDB: "password",
					HostDB:     "localhost",
					PortDB:     "54320",
					DBName:     "organization_db",
					DelayTime:  1,
				},
				EmailDealerConfig: models.EmailDealerConfig{Host: "", Port: 0, Username: "", Password: ""},
				LogLevel:          "",
			},
		},
	}
	for _, test := range testTable {
		if test.name == "positive" {
			log.On("DebugMsg", "fail connect to DB, try again").Return()
			log.On("DebugMsg", "connection to DB is OK").Return()
			db, _ := NewClient(ctx, &test.cfg, log)
			require.NotNil(t, db, "")
			defer db.dbConnect.Close(ctx)
		} else {
			log.On("DebugMsg", "fail connect to DB, try again").Return()
			db, _ := NewClient(ctx, &test.cfg, log)
			require.Nil(t, db, "")
		}
	}
}

func TestDB_CreateOrganization(t *testing.T) {
	ctx := context.Background()
	log := mocks.NewLogger(t)
	log.On("DebugMsg", "connection to DB is OK").Return()
	db, _ := NewClient(ctx, &testCfg, log)
	defer db.dbConnect.Close(ctx)

	testTable := []struct {
		name             string
		testOrganization models.Organization
	}{
		{
			name: "positive",
			testOrganization: models.Organization{
				Name:         "test_organization1",
				LegalType:    "test_type",
				LegalAddress: "test_address",
				INN:          "123456789010",
				OwnerID:      11,
			},
		}, {
			name: "negative",
			testOrganization: models.Organization{
				Name:         " ",
				LegalType:    "test_type",
				LegalAddress: "test_address",
				INN:          "123456789010",
				OwnerID:      11,
			},
		}, {
			name: "negative",
			testOrganization: models.Organization{
				Name:         "test_organization1",
				LegalType:    "test_type",
				LegalAddress: "test_address",
				INN:          "1234567890101112",
				OwnerID:      11,
			},
		},
	}
	for _, test := range testTable {
		if test.name == "positive" {
			err := db.CreateOrganization(ctx, &test.testOrganization)
			assert.Nil(t, err, "")
		} else {
			err := db.CreateOrganization(ctx, &test.testOrganization)
			assert.NotNil(t, err, "")
		}
	}
}

func TestDB_EditOrganization(t *testing.T) {
	ctx := context.Background()
	log := mocks.NewLogger(t)
	log.On("DebugMsg", "connection to DB is OK").Return()
	db, _ := NewClient(ctx, &testCfg, log)
	defer db.dbConnect.Close(ctx)

	testTable := []struct {
		name             string
		testOrganization models.Organization
	}{
		{
			name: "positive",
			testOrganization: models.Organization{
				Name:         "test_organization1_after edit",
				LegalType:    "test_type",
				LegalAddress: "test_address",
				INN:          "123456789010",
				OwnerID:      11,
			},
		}, {
			name: "negative",
			testOrganization: models.Organization{
				Name:         "test_organization1_after edit",
				LegalType:    "test_type",
				LegalAddress: "test_address",
				INN:          "123456789010",
				OwnerID:      12,
			},
		},
	}
	for _, test := range testTable {
		if test.name == "positive" {
			err := db.EditOrganization(ctx, &test.testOrganization)
			assert.Nil(t, err, "")
		} else {
			err := db.EditOrganization(ctx, &test.testOrganization)
			assert.NotNil(t, err, "")
		}
	}
}

func TestDB_GetInfoOrganization(t *testing.T) {
	ctx := context.Background()
	log := mocks.NewLogger(t)
	log.On("DebugMsg", "connection to DB is OK").Return()
	db, _ := NewClient(ctx, &testCfg, log)
	defer db.dbConnect.Close(ctx)
	testTable := []struct {
		name                 string
		orgID                int
		ExpectedOrganization models.Organization
	}{
		{
			name:  "positive",
			orgID: 2,
			ExpectedOrganization: models.Organization{
				OrganizationId: 2,
				Name:           "test_organization1_after edit",
				LegalType:      "test_type",
				LegalAddress:   "test_address",
				INN:            "123456789010",
				OwnerID:        11,
			},
		}, {
			name:                 "negative",
			orgID:                10,
			ExpectedOrganization: models.Organization{},
		},
	}

	for _, test := range testTable {
		if test.name == "positive" {
			result, err := db.GetInfoOrganization(ctx, test.orgID)
			assert.Nil(t, err, "")
			assert.Equal(t, test.ExpectedOrganization, *result)
		} else {
			result, err := db.GetInfoOrganization(ctx, test.orgID)
			assert.NotNil(t, err, "")
			assert.Nil(t, result, "")
		}
	}
}

func TestDB_CreateFilial(t *testing.T) {
	ctx := context.Background()
	log := mocks.NewLogger(t)
	log.On("DebugMsg", "connection to DB is OK").Return()
	db, _ := NewClient(ctx, &testCfg, log)
	defer db.dbConnect.Close(ctx)

	testTable := []struct {
		name      string
		filialDTO models.FilialDTO
	}{
		{
			name: "positive",
			filialDTO: models.FilialDTO{
				UserID:         10,
				FilialName:     "Filial A",
				Country:        "United States",
				City:           "New York",
				Address:        "123 Main Street",
				FilialType:     "Retail",
				PhoneList:      "+1-123-456-7890",
				EmailList:      "filial@example.com",
				PhotoIDList:    "photo-1,photo-2,photo-3",
				OrganizationID: 1,
			},
		}, {
			name: "positive",
			filialDTO: models.FilialDTO{
				UserID:         1,
				FilialName:     "Filial B",
				Country:        "Canada",
				City:           "Toronto",
				Address:        "456 Elm Street",
				FilialType:     "Office",
				PhoneList:      "+1-987-654-3210",
				EmailList:      "another.filial@example.com",
				PhotoIDList:    "photo-4,photo-5,photo-6",
				OrganizationID: 1,
				DirectorID:     1,
			},
		},
	}

	for _, test := range testTable {
		if test.name == "positive" {
			err := db.CreateFilial(ctx, &test.filialDTO)
			assert.Nil(t, err, "")
		} else {
			err := db.CreateFilial(ctx, &test.filialDTO)
			assert.NotNil(t, err, "")
		}
	}

}

func TestDB_EditFilial(t *testing.T) {
	ctx := context.Background()
	log := mocks.NewLogger(t)
	log.On("DebugMsg", "connection to DB is OK").Return()
	db, _ := NewClient(ctx, &testCfg, log)
	defer db.dbConnect.Close(ctx)

	testTable := []struct {
		name      string
		filialDTO models.FilialDTO
	}{
		{
			name: "positive",
			filialDTO: models.FilialDTO{
				UserID:         1,
				FilialID:       1,
				FilialName:     "Filial A after Edit",
				Country:        "United States",
				City:           "New York",
				Address:        "123 Main Street",
				FilialType:     "Retail",
				PhoneList:      "+1-123-456-7890",
				EmailList:      "filial@example.com",
				PhotoIDList:    "photo-1,photo-2,photo-3",
				OrganizationID: 1,
			},
		}, {
			name: "positive",
			filialDTO: models.FilialDTO{
				UserID:         10,
				FilialID:       2,
				FilialName:     "Filial B after Edit",
				Country:        "Canada",
				City:           "Toronto",
				Address:        "456 Elm Street",
				FilialType:     "Office",
				PhoneList:      "+1-987-654-3210",
				EmailList:      "another.filial@example.com",
				PhotoIDList:    "photo-4,photo-5,photo-6",
				OrganizationID: 1,
			},
		}, {
			name: "negative",
			filialDTO: models.FilialDTO{
				UserID:         11,
				FilialID:       2,
				FilialName:     "Filial B after Edit with negative",
				Country:        "Canada",
				City:           "Toronto",
				Address:        "456 Elm Street",
				FilialType:     "Office",
				PhoneList:      "+1-987-654-3210",
				EmailList:      "another.filial@example.com",
				PhotoIDList:    "photo-4,photo-5,photo-6",
				OrganizationID: 1,
			},
		},
	}

	for _, test := range testTable {
		if test.name == "positive" {
			err := db.EditFilial(ctx, &test.filialDTO)
			assert.Nil(t, err, "")
		} else {
			err := db.EditFilial(ctx, &test.filialDTO)
			assert.NotNil(t, err, "")
		}
	}
}

func TestDB_GetInfoFilial(t *testing.T) {
	ctx := context.Background()
	log := mocks.NewLogger(t)
	log.On("DebugMsg", "connection to DB is OK").Return()
	db, _ := NewClient(ctx, &testCfg, log)
	defer db.dbConnect.Close(ctx)
	testTable := []struct {
		name           string
		filialID       int
		ExpectedFilial models.Filial
	}{
		{
			name:     "positive",
			filialID: 1,
			ExpectedFilial: models.Filial{
				FilialID:       1,
				FilialName:     "Filial A after Edit",
				Country:        "United States",
				City:           "New York",
				Address:        "123 Main Street",
				FilialType:     "Retail",
				PhoneList:      "+1-123-456-7890",
				EmailList:      "filial@example.com",
				PhotoIDList:    "photo-1,photo-2,photo-3",
				OrganizationID: 1,
				DirectorID:     0,
			},
		}, {
			name:     "positive",
			filialID: 2,
			ExpectedFilial: models.Filial{
				FilialID:       2,
				FilialName:     "Filial B after Edit",
				Country:        "Canada",
				City:           "Toronto",
				Address:        "456 Elm Street",
				FilialType:     "Office",
				PhoneList:      "+1-987-654-3210",
				EmailList:      "another.filial@example.com",
				PhotoIDList:    "photo-4,photo-5,photo-6",
				OrganizationID: 1,
				DirectorID:     1,
			},
		},
	}

	for _, test := range testTable {
		if test.name == "positive" {
			result, err := db.GetInfoFilial(ctx, test.filialID)
			assert.Nil(t, err, "")
			assert.Equal(t, test.ExpectedFilial, *result)
		} else {
			result, err := db.GetInfoFilial(ctx, test.filialID)
			assert.NotNil(t, err, "")
			assert.Nil(t, result, "")
		}
	}
}

func TestDB_GetOrganizationFilials(t *testing.T) {
	ctx := context.Background()
	log := mocks.NewLogger(t)
	log.On("DebugMsg", "connection to DB is OK").Return()
	db, _ := NewClient(ctx, &testCfg, log)
	defer db.dbConnect.Close(ctx)

	testTable := []struct {
		name         string
		orgID        int
		ExpectedData []*models.OrganizationFilialsDTO
	}{
		{
			name:  "positive",
			orgID: 1,
			ExpectedData: []*models.OrganizationFilialsDTO{
				{
					FilialID:   1,
					FilialName: "Filial A after Edit",
				}, {
					FilialID:   2,
					FilialName: "Filial B after Edit",
				},
			},
		}, {
			name:         "positive",
			orgID:        2,
			ExpectedData: []*models.OrganizationFilialsDTO{},
		}, {
			name:         "negative",
			orgID:        5,
			ExpectedData: []*models.OrganizationFilialsDTO{},
		},
	}
	for _, test := range testTable {
		if test.name == "positive" {
			result, err := db.GetOrganizationFilials(ctx, test.orgID)
			assert.Nil(t, err, "")
			assert.Equal(t, test.ExpectedData, result)
		} else {
			result, err := db.GetOrganizationFilials(ctx, test.orgID)
			assert.Nil(t, err, "")
			assert.Equal(t, test.ExpectedData, result)
		}
	}
}

func TestDB_AddWorkerFilial(t *testing.T) {
	ctx := context.Background()
	log := mocks.NewLogger(t)
	log.On("DebugMsg", "connection to DB is OK").Return()
	db, _ := NewClient(ctx, &testCfg, log)
	defer db.dbConnect.Close(ctx)

	testTable := []struct {
		name       string
		token      string
		UserData   models.AddWorkersDTO
		ExpectedID int
	}{
		{
			name:  "positive",
			token: "e9acbc6811001fd31105cf526e1f523d",
			UserData: models.AddWorkersDTO{
				OrganizationID: 1,
				FilialID:       2,
				Position:       "manager",
				Email:          "testUser1@mail.com",
			},
			ExpectedID: 2,
		}, {
			name:  "positive",
			token: "e9acbc6811001fd31105cf526e1s123a",
			UserData: models.AddWorkersDTO{
				OrganizationID: 2,
				Position:       "manager",
				Email:          "testUser2@mail.com",
			},
			ExpectedID: 3,
		}, {
			name:  "negative",
			token: "e9acbc6811001fd31105cf526e1s123a",
			UserData: models.AddWorkersDTO{
				OrganizationID: 5,
				FilialID:       3,
				Position:       "manager",
				Email:          "testUser2@mail.com",
			},
			ExpectedID: 0,
		},
	}
	for _, test := range testTable {
		if test.name == "positive" {
			Id, err := db.AddWorkerFilial(ctx, test.token, &test.UserData)
			assert.Nil(t, err, "")
			assert.Equal(t, test.ExpectedID, Id)
		} else {
			Id, err := db.AddWorkerFilial(ctx, test.token, &test.UserData)
			assert.NotNil(t, err, "")
			assert.Equal(t, test.ExpectedID, Id)
		}
	}
}

func TestDB_GetEmployeeInfoByUserID(t *testing.T) {
	ctx := context.Background()
	log := mocks.NewLogger(t)
	log.On("DebugMsg", "connection to DB is OK").Return()
	db, _ := NewClient(ctx, &testCfg, log)
	defer db.dbConnect.Close(ctx)

	testTable := []struct {
		name             string
		userId           int
		expectedEmployee models.Employee
	}{
		{
			name:   "positive",
			userId: 2,
			expectedEmployee: models.Employee{
				ID:                    2,
				OrganizationID:        1,
				FilialID:              2,
				Position:              "manager",
				Email:                 "testUser1@mail.com",
				EmailConfirmationFlag: false,
			},
		}, {
			name:             "negative",
			userId:           70,
			expectedEmployee: models.Employee{},
		},
	}
	for _, test := range testTable {
		if test.name == "positive" {
			result, err := db.GetEmployeeInfoByUserID(ctx, test.userId)
			assert.Nil(t, err, "")
			assert.Equal(t, test.expectedEmployee, *result)
		} else {
			result, err := db.GetEmployeeInfoByUserID(ctx, test.userId)
			assert.NotNil(t, err, "")
			assert.Nil(t, result)
		}
	}
}
