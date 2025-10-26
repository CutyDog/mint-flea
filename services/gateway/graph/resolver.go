package graph

import (
	"github.com/CutyDog/mint-flea/services/gateway/client"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AccountClient *client.AccountClient
	WalletClient  *client.WalletClient
}
