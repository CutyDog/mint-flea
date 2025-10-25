package auth

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

type FirebaseAuth struct {
	client *auth.Client
}

// Firebase Admin SDKの初期化
func NewFirebaseApp(ctx context.Context) (*firebase.App, error) {
	config := LoadFirebaseConfig()

	var opts []option.ClientOption

	if config.ServiceAccount != "" {
		// サービスアカウントキーをJSON文字列として読み込み
		opts = append(opts, option.WithCredentialsJSON([]byte(config.ServiceAccount)))
	}

	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: config.ProjectID,
	}, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize firebase app: %w", err)
	}

	return app, nil
}

// Firebase Authの初期化
func NewFirebaseAuth(ctx context.Context) (*FirebaseAuth, error) {
	app, err := NewFirebaseApp(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize firebase app: %w", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get firebase auth client: %w", err)
	}

	return &FirebaseAuth{
		client: client,
	}, nil
}
