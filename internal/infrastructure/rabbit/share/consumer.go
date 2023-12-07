package shareRabbit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"sbp/internal/app"
	"sbp/internal/conversions"
	"sbp/internal/entities"
	rabbitEntities "sbp/internal/entities/rabbit"
	"sbp/pkg/rabbitmq"

	gorabbit "github.com/wagslane/go-rabbitmq"

	"go.uber.org/zap"
)

type shareClient struct {
	logger          *zap.SugaredLogger
	startupConsumer *rabbitmq.Consumer
	adminConsumer   *rabbitmq.Consumer
	publisher       app.SharePublisher
}

type shareService interface {
	UpsertUser(ctx context.Context, user entities.User) error
	UpsertGroup(ctx context.Context, group entities.Group) error
	UpsertOrganization(ctx context.Context, org entities.Organization) error
}

const MaxRetries = 10

func NewShareConsumer(logger *zap.SugaredLogger, client *rabbitmq.RabbitMqClient, share shareService, publisher app.SharePublisher) (*shareClient, error) {
	if client == nil {
		return nil, errors.New("NewShareConsumer: client = nil")
	}

	handler, err := createHandler(logger, share, publisher)
	if err != nil {
		return nil, err
	}

	shareCon, err := rabbitmq.NewConsumer(logger,
		client,
		handler,
		string(rabbitEntities.SBPAdminDataQueue),

		gorabbit.WithConsumerOptionsExchangeDeclare,

		gorabbit.WithConsumerOptionsExchangeName(string(rabbitEntities.WashBonusExchange)),
		gorabbit.WithConsumerOptionsExchangeKind("direct"),

		gorabbit.WithConsumerOptionsRoutingKey(string(rabbitEntities.WashBonusRoutingKey)),
		gorabbit.WithConsumerOptionsExchangeDurable,
	)
	if err != nil {
		return nil, err
	}

	adminCon, err := rabbitmq.NewConsumer(logger,
		client,
		handler,
		string(rabbitEntities.SBPStartupQueue),

		gorabbit.WithConsumerOptionsExchangeDeclare,

		gorabbit.WithConsumerOptionsExchangeName(string(rabbitEntities.AdminsExchange)),
		gorabbit.WithConsumerOptionsExchangeKind("fanout"),

		gorabbit.WithConsumerOptionsRoutingKey(string(rabbitEntities.SBPStartupQueue)),
		gorabbit.WithConsumerOptionsExchangeDurable,
	)
	if err != nil {
		return nil, err
	}

	return &shareClient{
		logger:          logger,
		startupConsumer: shareCon,
		adminConsumer:   adminCon,
		publisher:       publisher,
	}, nil
}

func (c shareClient) Close() {
	c.startupConsumer.Close()
	c.adminConsumer.Close()
}

func getRetryCount(headers map[string]interface{}) int {
	if count, ok := headers["x-retry-count"].(int); ok {
		return count
	}
	return 0
}

func createHandler(logger *zap.SugaredLogger, share shareService, publisher app.SharePublisher) (rabbitmq.ConsumerHandler, error) {
	return func(ctx context.Context, d rabbitmq.RbqMessage) error {
		messageType := rabbitEntities.Message(d.Type)

		retryCount := getRetryCount(d.Headers)
		retryCount++

		if retryCount > MaxRetries {
			return fmt.Errorf("Max retries reached, handling failure")
		}

		switch messageType {
		case rabbitEntities.OrganizationMessage:
			var msg rabbitEntities.Organization
			err := json.Unmarshal(d.Body, &msg)
			if err != nil {
				return err
			}

			org, err := conversions.OrganizationFromRabbit(msg)
			if err != nil {
				return fmt.Errorf("unable to map organization from rabbit: %w", err)
			}

			err = share.UpsertOrganization(ctx, org)
			if errors.Is(err, entities.ErrBadVersion) {
				return err
			}

			if err != nil {
				d.Nack(false, true)
				return err
			}

		case rabbitEntities.ServerGroupMessage:
			var msg rabbitEntities.ServerGroup
			err := json.Unmarshal(d.Body, &msg)
			if err != nil {
				return err
			}

			group, err := conversions.GroupFromRabbit(msg)
			if err != nil {
				return fmt.Errorf("unable to map group from rabbit: %w", err)
			}

			err = share.UpsertGroup(ctx, group)
			if errors.Is(err, entities.ErrBadVersion) {
				return err
			}

			if err != nil {
				d.Nack(false, true)
				return err
			}

		case rabbitEntities.AdminUserMessage:
			var msg rabbitEntities.AdminUser
			err := json.Unmarshal(d.Body, &msg)
			if err != nil {
				return err
			}

			user, err := conversions.UserFromRabbit(msg)
			if err != nil {
				return fmt.Errorf("unable to map user from rabbit: %w", err)
			}

			err = share.UpsertUser(ctx, user)
			if errors.Is(err, entities.ErrBadVersion) {
				return err
			}

			if err != nil {
				d.Nack(false, true)
				return err
			}

		default:
			return fmt.Errorf("received unexpected message with type: %s", messageType)
		}

		return nil
	}, nil
}
