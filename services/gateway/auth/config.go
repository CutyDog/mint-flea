package auth

import (
	"os"
)

type FirebaseConfig struct {
	ProjectID      string
	ServiceAccount string
}

// LoadFirebaseConfig loads Firebase configuration from environment variables
func LoadFirebaseConfig() *FirebaseConfig {
	serviceAccountPath := os.Getenv("FIREBASE_SERVICE_ACCOUNT_PATH")

	var serviceAccount string
	if serviceAccountPath != "" {
		// ファイルからサービスアカウントキーを読み込み
		if data, err := os.ReadFile(serviceAccountPath); err == nil {
			serviceAccount = string(data)
		}
	}

	return &FirebaseConfig{
		ProjectID:      os.Getenv("FIREBASE_PROJECT_ID"),
		ServiceAccount: serviceAccount,
	}
}
