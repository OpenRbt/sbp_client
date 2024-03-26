package repository_test

import (
	"context"
	"fmt"
	"sbp/internal/config"
	"sbp/internal/entities"
	"sbp/internal/repository"
	"sbp/internal/testutils"
	"testing"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

func TestTransactionRepository_List(t *testing.T) {
	port := testutils.SetupDatabase(t)
	defer testutils.TeardownDatabase(t)

	logger, _ := zap.NewDevelopment()
	repo, err := repository.NewRepository(config.RepositoryConfig{
		DBConfig: &config.DBConfig{
			Host:           "localhost",
			Port:           fmt.Sprintf("%d", port),
			Database:       "testdb",
			User:           "postgres",
			Password:       "secret",
			MigrationsPath: "../repository/migrations",
		},
		Logger: logger.Sugar(),
	})
	if err != nil {
		t.Fatalf("Failed to create user repo: %s", err)
	}

	testOrg := entities.Organization{
		ID:          uuid.NewV4(),
		Name:        "default",
		DisplayName: "default",
		Description: "default",
		IsDefault:   true,
		Deleted:     false,
		Version:     0,
	}

	testGroup := entities.Group{
		ID:             uuid.NewV4(),
		OrganizationID: testOrg.ID,
		Name:           "default",
		Description:    "default",
		IsDefault:      true,
		Deleted:        false,
		Version:        0,
	}

	testUser := entities.User{
		ID:      "test",
		Email:   "test@mail.ru",
		Name:    "test",
		Role:    "system_manager",
		Version: 0,
		Deleted: false,
	}

	testWash := entities.WashCreation{
		OwnerID:          testUser.ID,
		Password:         "test",
		Title:            "test",
		Description:      "test",
		TerminalKey:      "test",
		TerminalPassword: "test",
		GroupID:          testGroup.ID,
	}

	err = repo.InsertUser(context.Background(), testUser)
	if err != nil {
		t.Fatalf("Failed to insert user: %s", err)
	}

	err = repo.InsertOrganization(context.Background(), testOrg)
	if err != nil {
		t.Fatalf("Failed to insert default organization: %s", err)
	}

	err = repo.InsertGroup(context.Background(), testGroup)
	if err != nil {
		t.Fatalf("Failed to insert test group: %v", err)
	}

	wash, err := repo.CreateWash(context.Background(), testWash)
	if err != nil {
		t.Fatalf("Failed to create test wash: %v", err)
	}

	testTransaction := entities.TransactionCreate{
		ID:            uuid.NewV4(),
		WashID:        wash.ID.String(),
		PostID:        "1",
		Amount:        100,
		PaymentIDBank: "1",
		Status:        entities.TransactionStatusAuthorized,
	}

	err = repo.CreateTransaction(context.Background(), testTransaction)
	if err != nil {
		t.Fatalf("Failed to create transaction: %v", err)
	}

	list, err := repo.TransactionsList(context.Background(), entities.TransactionFilter{
		Filter: entities.NewFilter(1, 10),
	})
	if err != nil {
		t.Fatalf("Failed to get transactions list: %v", err)
	}
	if len(list.Items) != 1 {
		t.Fatalf("Failed to get transactions list: len if list != 1")
	}
	if list.Page != 1 {
		t.Fatalf("Failed to get transactions list: page != 1")
	}
	if list.PageSize != 10 {
		t.Fatalf("Failed to get transactions list: pageSize != 10")
	}
	if list.TotalPages != 1 {
		t.Fatalf("Failed to get transactions list: totalPages != 1")
	}
	if list.TotalItems != 1 {
		t.Fatalf("Failed to get transactions list: totalItems != 1")
	}
	if !uuid.Equal(list.Items[0].ID, testTransaction.ID) {
		t.Fatalf("Failed to get transactions list: transaction.ID != testTransaction.ID")
	}
	if !uuid.Equal(list.Items[0].Wash.ID, wash.ID) {
		t.Fatalf("Failed to get transactions list: transaction.Wash.ID != testWash.ID")
	}
	if !uuid.Equal(list.Items[0].Group.ID, testGroup.ID) {
		t.Fatalf("Failed to get transactions list: transaction.Group.ID != testGroup.ID")
	}
	if !uuid.Equal(list.Items[0].Organization.ID, testOrg.ID) {
		t.Fatalf("Failed to get transactions list: transaction.Organization.ID != testOrg.ID")
	}
}
