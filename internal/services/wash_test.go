package services_test

import (
	"context"
	"fmt"
	"sbp/internal/config"
	"sbp/internal/entities"
	leawash "sbp/internal/infrastructure/rabbit/lea"
	"sbp/internal/repository"
	"sbp/internal/services"
	"sbp/internal/testutils"
	"testing"

	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

func TestWashService_AssignWashToGroup(t *testing.T) {
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

	broker, err := leawash.NewBrokerUserCreator(nil)
	if err != nil {
		t.Fatalf("Failed to create user broker: %s", err)
	}

	washService, err := services.NewWashService(context.Background(), logger.Sugar(), repo, broker, 10)
	if err != nil {
		t.Fatalf("Failed to create user service: %s", err)
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

	testCases := []struct {
		name        string
		auth        *entities.Auth
		washID      uuid.UUID
		groupID     uuid.UUID
		expectError bool
	}{
		{
			name:        "SuccessfulAssignment",
			auth:        &entities.Auth{User: entities.User{Role: entities.SystemManagerRole}},
			washID:      wash.ID,
			groupID:     testGroup.ID,
			expectError: false,
		},
		{
			name:        "AssignmentWithoutPermission",
			auth:        &entities.Auth{User: entities.User{Role: entities.NoAccessRole}},
			washID:      wash.ID,
			groupID:     testGroup.ID,
			expectError: true,
		},
		{
			name:        "NonExistentWash",
			auth:        &entities.Auth{User: entities.User{Role: entities.SystemManagerRole}},
			washID:      uuid.NewV4(),
			groupID:     testGroup.ID,
			expectError: true,
		},
		{
			name:        "NonExistentGroup",
			auth:        &entities.Auth{User: entities.User{Role: entities.SystemManagerRole}},
			washID:      wash.ID,
			groupID:     uuid.NewV4(),
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := washService.AssignWashToGroup(context.Background(), tc.auth, tc.washID, tc.groupID)

			if (err != nil) != tc.expectError {
				t.Errorf("Test %v: Expected error: %v, got: %v", tc.name, tc.expectError, err)
			}
		})
	}
}
