package services_test

import (
	"context"
	"fmt"
	"sbp/internal/config"
	"sbp/internal/entities"
	"sbp/internal/repository"
	"sbp/internal/services"
	"sbp/internal/testutils"
	"testing"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func TestUserService_UpsertUser(t *testing.T) {
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

	userService, err := services.NewUserService(context.Background(), logger.Sugar(), repo)
	if err != nil {
		t.Fatalf("Failed to create user service: %s", err)
	}

	testCases := []struct {
		name     string
		user     entities.User
		wantErr  bool
		setup    func()
		teardown func()
	}{
		{
			name: "Insert New User",
			user: entities.User{
				ID:      "new-user-id",
				Email:   "new@example.com",
				Name:    "New User",
				Role:    "no_access",
				Version: 1,
			},
			wantErr: false,
		},
		{
			name: "Update Existing User",
			user: entities.User{
				ID:      "existing-user-id",
				Email:   "existing@example.com",
				Name:    "Existing User",
				Role:    "admin",
				Version: 2,
			},
			wantErr: false,
			setup: func() {
				if err := userService.UpsertUser(context.Background(), entities.User{
					ID:      "existing-user-id",
					Email:   "existing@example.com",
					Name:    "Existing User",
					Role:    "no_access",
					Version: 1,
				}); err != nil {
					t.Errorf("Failed to insert existing user: %s", err.Error())
				}
			},
		},
		{
			name: "The same Version",
			user: entities.User{
				ID:      "bad-version",
				Email:   "badVersion@example.com",
				Name:    "Existing User",
				Role:    "no_access",
				Version: 3,
			},
			wantErr: true,
			setup: func() {
				if err := userService.UpsertUser(context.Background(), entities.User{
					ID:      "bad-version",
					Email:   "badVersion@example.com",
					Name:    "Existing User",
					Role:    "no_access",
					Version: 3,
				}); err != nil {
					t.Errorf("Failed to insert existing user: %s", err.Error())
				}
			},
		},
		{
			name: "Lower Version",
			user: entities.User{
				ID:      "lower-version",
				Email:   "lowerVersion@example.com",
				Name:    "Existing User",
				Role:    "no_access",
				Version: 2,
			},
			wantErr: true,
			setup: func() {
				if err := userService.UpsertUser(context.Background(), entities.User{
					ID:      "lower-version",
					Email:   "lowerVersion@example.com",
					Name:    "Existing User",
					Role:    "no_access",
					Version: 3,
				}); err != nil {
					t.Errorf("Failed to insert existing user: %s", err.Error())
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setup != nil {
				tc.setup()
			}

			err := userService.UpsertUser(context.Background(), tc.user)

			if (err != nil) != tc.wantErr {
				t.Errorf("UpsertUser() error = %v, wantErr %v", err, tc.wantErr)
			}

			if tc.teardown != nil {
				tc.teardown()
			}
		})
	}
}
