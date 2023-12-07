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

func TestGroupService_UpsertGroup(t *testing.T) {
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

	groupService, err := services.NewGroupService(context.Background(), logger.Sugar(), repo)
	if err != nil {
		t.Fatalf("Failed to create group service: %s", err)
	}

	defaultOrg := entities.Organization{
		ID:          uuid.NewV4(),
		Name:        "default",
		DisplayName: "default",
		Description: "default",
		IsDefault:   true,
		Deleted:     false,
		Version:     0,
	}

	err = orgService.UpsertOrganization(context.Background(), defaultOrg)
	if err != nil {
		t.Fatalf("Failed to insert default organization: %s", err)
	}

	testCases := []struct {
		name     string
		groupID  uuid.UUID
		group    entities.Group
		wantErr  bool
		setup    func(uuid.UUID)
		teardown func()
	}{
		{
			name:    "Insert New Group",
			groupID: uuid.NewV4(),
			group: entities.Group{
				OrganizationID: defaultOrg.ID,
				Name:           "New Group",
				Description:    "New",
				IsDefault:      false,
				Deleted:        false,
				Version:        0,
			},
			wantErr: false,
		},
		{
			name:    "Update Existing Group",
			groupID: uuid.NewV4(),
			group: entities.Group{
				OrganizationID: defaultOrg.ID,
				Name:           "Updated Group",
				Description:    "Updated",
				IsDefault:      false,
				Deleted:        false,
				Version:        1,
			},
			wantErr: false,
			setup: func(id uuid.UUID) {
				if err := groupService.UpsertGroup(context.Background(), entities.Group{
					ID:             id,
					OrganizationID: defaultOrg.ID,
					Name:           "Existing Group",
					Description:    "Existing",
					IsDefault:      false,
					Deleted:        false,
					Version:        0,
				}); err != nil {
					t.Errorf("Failed to insert existing group: %s", err.Error())
				}
			},
		},
		{
			name:    "The same Version",
			groupID: uuid.NewV4(),
			group: entities.Group{
				OrganizationID: defaultOrg.ID,
				Name:           "Bad version Group",
				Description:    "Bad",
				IsDefault:      false,
				Deleted:        false,
				Version:        1,
			},
			wantErr: true,
			setup: func(id uuid.UUID) {
				if err := groupService.UpsertGroup(context.Background(), entities.Group{
					ID:             id,
					OrganizationID: defaultOrg.ID,
					Name:           "Existing Group",
					Description:    "Existing",
					IsDefault:      false,
					Deleted:        false,
					Version:        1,
				}); err != nil {
					t.Errorf("Failed to insert existing group: %s", err.Error())
				}
			},
		},
		{
			name:    "Lower Version",
			groupID: uuid.NewV4(),
			group: entities.Group{
				OrganizationID: defaultOrg.ID,
				Name:           "Lower version Group",
				Description:    "Lower",
				IsDefault:      false,
				Deleted:        false,
				Version:        0,
			},
			wantErr: true,
			setup: func(id uuid.UUID) {
				if err := groupService.UpsertGroup(context.Background(), entities.Group{
					ID:             id,
					OrganizationID: defaultOrg.ID,
					Name:           "Existing Group",
					Description:    "Existing",
					IsDefault:      false,
					Deleted:        false,
					Version:        1,
				}); err != nil {
					t.Errorf("Failed to insert existing group: %s", err.Error())
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.group.ID = tc.groupID

			if tc.setup != nil {
				tc.setup(tc.groupID)
			}

			err := groupService.UpsertGroup(context.Background(), tc.group)

			if (err != nil) != tc.wantErr {
				t.Errorf("UpsertGroup() error = %v, wantErr %v", err, tc.wantErr)
			}

			if tc.teardown != nil {
				tc.teardown()
			}
		})
	}
}
