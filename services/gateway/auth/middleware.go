package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/CutyDog/mint-flea/services/gateway/errors"
)

type AuthMiddleware struct {
	firebaseAuth *FirebaseAuth
}

type AuthContextKey string

const (
	UserContextKey AuthContextKey = "user"
)

func NewAuthMiddleware() (*AuthMiddleware, error) {
	// Firebase Authを初期化
	firebaseAuth, err := NewFirebaseAuth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize firebase auth: %w", err)
	}

	return &AuthMiddleware{
		firebaseAuth: firebaseAuth,
	}, nil
}

func (am *AuthMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// GraphQLのPOSTリクエストのみを処理
		if r.Method != "POST" {
			next.ServeHTTP(w, r)
			return
		}

		// 認証処理
		am.handleAuth(r, w, next)
	})
}

// Authorization ヘッダーから JWT を取得して検証、コンテキストにユーザー情報を保存して次のハンドラーに渡す
func (am *AuthMiddleware) handleAuth(r *http.Request, w http.ResponseWriter, next http.Handler) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		// 認証ヘッダーがない場合は認証エラーを返す
		log.Printf("No authorization header found, authentication required")
		errors.SendUnauthenticatedError(w, "Authentication required")
		return
	}

	// JWTを検証
	token, err := am.VerifyIDToken(r.Context(), authHeader)
	if err != nil {
		// JWTが無効な場合は認証エラーを返す
		log.Printf("Unauthorized: Invalid token: %v", err)
		errors.SendUnauthenticatedError(w, "Invalid token")
		return
	}

	log.Printf("Successfully authenticated user: %s", token.UID)
	// コンテキストにユーザー情報を保存
	ctx := context.WithValue(r.Context(), UserContextKey, token)
	next.ServeHTTP(w, r.WithContext(ctx))
}

// トークンを検証してユーザー情報を取得
func (am *AuthMiddleware) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	idToken, found := strings.CutPrefix(idToken, "Bearer ")
	if !found {
		return nil, fmt.Errorf("invalid ID token")
	}

	token, err := am.firebaseAuth.client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, fmt.Errorf("failed to verify ID token: %w", err)
	}

	return token, nil
}
