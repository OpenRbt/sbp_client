package firebase_authorization

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	logicEntities "sbp/internal/logic/entities"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
)

const authTimeout = time.Second

// FirebaseClient ...
type FirebaseClient struct {
	app    *firebase.App
	client *auth.Client
}

// NewAuthClient ...
func NewAuthClient(keyfileLocation string) (*FirebaseClient, error) {
	keyFilePath, err := filepath.Abs(keyfileLocation)
	if err != nil {
		return nil, errors.New("Unable to load Client key")
	}
	opt := option.WithCredentialsFile(keyFilePath)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, errors.New("Failed to load Firebase")
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, errors.New("Failed to load Firebase auth")
	}

	return &FirebaseClient{
		app:    app,
		client: client,
	}, nil
}

// Auth ...
func (firebaseClient *FirebaseClient) Auth(bearer string) (*logicEntities.Auth, error) {

	// parse token id
	idToken := strings.TrimSpace(strings.Replace(bearer, "Bearer", "", 1))

	if idToken == "" {
		return nil, logicEntities.ErrAccessDenied
	}

	// context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), authTimeout)
	defer cancel()

	// verify token id
	token, err := firebaseClient.client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, logicEntities.ErrAccessDenied
	}

	// get user
	user, err := firebaseClient.client.GetUser(ctx, token.UID)
	if err != nil {
		if firebaseErr, ok := err.(*googleapi.Error); ok {
			if firebaseErr.Code == 401 {
				return nil, logicEntities.ErrAccessDenied
			} else if firebaseErr.Code == 403 {
				return nil, logicEntities.ErrAccessDenied
			}
		}
		return nil, logicEntities.ErrAccessDenied
	}

	return &logicEntities.Auth{
		UID:          user.UID,
		Disabled:     user.Disabled,
		UserMetadata: (*logicEntities.AuthUserMeta)(user.UserMetadata),
	}, nil
}

// SignUP ... for tests TODO: delete before service release
func (firebaseClient *FirebaseClient) SignUP() (*logicEntities.Token, error) {
	client, err := firebaseClient.app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	ctx := context.Background()
	userId := "vtv8NDpeoZbnKPlleNHDJqdkg9f1"
	customToken, err := client.CustomToken(ctx, userId)
	if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
	}

	idToken, err := getIdToken(customToken)
	if err != nil {
		log.Fatalf("error minting get id token: %v\n", err)
	}

	return &logicEntities.Token{
		Value: idToken,
	}, nil
}

func getIdToken(customToken string) (string, error) {
	host := "https://identitytoolkit.googleapis.com"
	path := "v1/accounts:signInWithCustomToken"
	apiKey := "AIzaSyCzBqNuINYIU-eCeVJ_QtHqrafG83WAGqY"
	url := fmt.Sprintf("%s/%s?key=%s", host, path, apiKey)
	method := "POST"
	s := fmt.Sprintf(`{
		"token": "%s",
		"returnSecureToken": true
	}`, customToken)
	payload := strings.NewReader(s)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")

	resp := struct {
		Kind         string `json:"kind"`
		IDToken      string `json:"idToken"`
		RefreshToken string `json:"refreshToken"`
		ExpiresIn    string `json:"expiresIn"`
		IsNewUser    bool   `json:"isNewUser"`
	}{}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return "", err
	}

	return resp.IDToken, nil
}
