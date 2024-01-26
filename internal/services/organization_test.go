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
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

func TestOrganizationService_UpsertOrganization(t *testing.T) {
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

	orgService, err := services.NewOrganizationService(context.Background(), logger.Sugar(), repo)
	if err != nil {
		t.Fatalf("Failed to create organization service: %s", err)
	}

	testCases := []struct {
		name     string
		orgID    uuid.UUID
		org      entities.Organization
		wantErr  bool
		setup    func(uuid.UUID)
		teardown func()
	}{
		{
			name:  "Insert New Organization",
			orgID: uuid.NewV4(),
			org: entities.Organization{
				Name:        "New Org",
				Description: "New",
				IsDefault:   false,
				Deleted:     false,
				Version:     0,
			},
			wantErr: false,
		},
		{
			name:  "Update Existing Organization",
			orgID: uuid.NewV4(),
			org: entities.Organization{
				Name:        "Updated Org",
				Description: "Updated",
				IsDefault:   false,
				Deleted:     false,
				Version:     1,
			},
			wantErr: false,
			setup: func(id uuid.UUID) {
				if err := orgService.UpsertOrganization(context.Background(), entities.Organization{
					ID:          id,
					Name:        "Existing Org",
					Description: "Existing",
					IsDefault:   false,
					Deleted:     false,
					Version:     0,
				}); err != nil {
					t.Errorf("Failed to insert existing org: %s", err.Error())
				}
			},
		},
		{
			name:  "The same Version",
			orgID: uuid.NewV4(),
			org: entities.Organization{
				Name:        "Bad version Org",
				Description: "Bad",
				IsDefault:   false,
				Deleted:     false,
				Version:     1,
			},
			wantErr: true,
			setup: func(id uuid.UUID) {
				if err := orgService.UpsertOrganization(context.Background(), entities.Organization{
					ID:          id,
					Name:        "Existing Org",
					Description: "Existing",
					IsDefault:   false,
					Deleted:     false,
					Version:     1,
				}); err != nil {
					t.Errorf("Failed to insert existing org: %s", err.Error())
				}
			},
		},
		{
			name:  "Lower Version",
			orgID: uuid.NewV4(),
			org: entities.Organization{
				Name:        "Lower version Org",
				Description: "Lower",
				IsDefault:   false,
				Deleted:     false,
				Version:     0,
			},
			wantErr: true,
			setup: func(id uuid.UUID) {
				if err := orgService.UpsertOrganization(context.Background(), entities.Organization{
					ID:          id,
					Name:        "Existing Group",
					Description: "Existing",
					IsDefault:   false,
					Deleted:     false,
					Version:     1,
				}); err != nil {
					t.Errorf("Failed to insert existing org: %s", err.Error())
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.org.ID = tc.orgID

			if tc.setup != nil {
				tc.setup(tc.orgID)
			}

			err := orgService.UpsertOrganization(context.Background(), tc.org)

			if (err != nil) != tc.wantErr {
				t.Errorf("UpsertOrganization() error = %v, wantErr %v", err, tc.wantErr)
			}

			if tc.teardown != nil {
				tc.teardown()
			}
		})
	}
}
