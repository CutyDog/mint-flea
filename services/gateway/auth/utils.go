package auth

import (
	"context"
	"fmt"

	firebaseAuth "firebase.google.com/go/v4/auth"
)

// GetUserFromContext extracts user info from context
func GetUserFromContext(ctx context.Context) (*firebaseAuth.Token, error) {
	user, ok := ctx.Value(UserContextKey).(*firebaseAuth.Token)
	if !ok {
		return nil, fmt.Errorf("user not found in context")
	}
	return user, nil
}

// GetUserUIDFromContext extracts user UID from context
func GetUserUIDFromContext(ctx context.Context) (string, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return "", err
	}
	return user.UID, nil
}

// IsAuthenticated checks if the user is authenticated
func IsAuthenticated(ctx context.Context) bool {
	_, err := GetUserFromContext(ctx)
	return err == nil
}
