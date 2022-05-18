package adapters

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"fmt"
	"google.golang.org/api/option"
)

func NewFirebaseAuthClient(credentials []byte) (*auth.Client, error) {
	opt := option.WithCredentialsJSON(credentials)
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("NewApp: %w", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}

	return client, nil
}
