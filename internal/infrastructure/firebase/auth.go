package firebase

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"sbp/internal/entities"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"

	opErrors "github.com/go-openapi/errors"
)

const authTimeout = time.Second * 15

type UserRepo interface {
	GetUserByID(ctx context.Context, id string) (entities.User, error)
}

type FirebaseClient struct {
	app    *firebase.App
	client *auth.Client

	userRepo UserRepo
}

var ErrUnauthorized = opErrors.New(401, "unauthorized")

func NewAuthClient(keyfileLocation string, userRepo UserRepo) (*FirebaseClient, error) {
	keyFilePath, err := filepath.Abs(keyfileLocation)
	if err != nil {
		return nil, fmt.Errorf("Unable to load Client key: %w", err)
	}
	opt := option.WithCredentialsFile(keyFilePath)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("Failed to load Firebase: %w", err)
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Failed to load Firebase auth: %w", err)
	}

	return &FirebaseClient{
		app:    app,
		client: client,

		userRepo: userRepo,
	}, nil
}

func (fb *FirebaseClient) Auth(bearer string) (*entities.Auth, error) {
	ctx, cancel := context.WithTimeout(context.Background(), authTimeout)
	defer cancel()

	jwtToken := strings.TrimSpace(strings.Replace(bearer, "Bearer", "", 1))
	if jwtToken == "" {
		return nil, ErrUnauthorized
	}

	token, err := fb.client.VerifyIDToken(context.Background(), jwtToken)
	if err != nil {
		return nil, ErrUnauthorized
	}

	fbUser, err := fb.client.GetUser(ctx, token.UID)
	if err != nil {
		return nil, ErrUnauthorized
	}

	user, err := fb.userRepo.GetUserByID(ctx, token.UID)
	if err != nil {
		if errors.Is(err, entities.ErrNotFound) {
			err = ErrUnauthorized
		}

		return nil, err
	}

	return &entities.Auth{
		User:         user,
		UserMetadata: (*entities.UserAuthMetadata)(fbUser.UserMetadata),
		Disabled:     fbUser.Disabled,
	}, nil
}
