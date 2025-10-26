package client

import (
	"context"
	"fmt"

	accountv1 "github.com/CutyDog/mint-flea/proto/gen/account/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type WalletClient struct {
	client accountv1.WalletServiceClient
	conn   *grpc.ClientConn
}

func NewWalletClient(addr string) (*WalletClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to wallet service: %w", err)
	}

	client := accountv1.NewWalletServiceClient(conn)

	return &WalletClient{
		client: client,
		conn:   conn,
	}, nil
}

func (c *WalletClient) Close() error {
	return c.conn.Close()
}

func (c *WalletClient) ListWallets(ctx context.Context, accountID int64) ([]*accountv1.Wallet, error) {
	resp, err := c.client.ListWallets(ctx, &accountv1.ListWalletsRequest{
		AccountId: accountID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list wallets: %w", err)
	}
	return resp.Wallets, nil
}

func (c *WalletClient) LinkWallet(ctx context.Context, accountID int64, address string, chainId int64, isMain bool) (*accountv1.Wallet, error) {
	resp, err := c.client.LinkWallet(ctx, &accountv1.LinkWalletRequest{
		AccountId: accountID,
		Address:   address,
		ChainId:   chainId,
		IsMain:    isMain,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to link wallet: %w", err)
	}
	return resp.Wallet, nil
}

func (c *WalletClient) UnlinkWallet(ctx context.Context, accountID int64, walletID int64) (bool, error) {
	resp, err := c.client.UnlinkWallet(ctx, &accountv1.UnlinkWalletRequest{
		AccountId: accountID,
		WalletId:  walletID,
	})
	if err != nil {
		return false, fmt.Errorf("failed to unlink wallet: %w", err)
	}
	return resp.Success, nil
}

func (c *WalletClient) SetMainWallet(ctx context.Context, accountID int64, walletID int64) (*accountv1.Wallet, error) {
	resp, err := c.client.SetMainWallet(ctx, &accountv1.SetMainWalletRequest{
		AccountId: accountID,
		WalletId:  walletID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to set main wallet: %w", err)
	}
	return resp.Wallet, nil
}

func (c *WalletClient) GetMainWallet(ctx context.Context, accountID int64) (*accountv1.Wallet, error) {
	resp, err := c.client.GetMainWallet(ctx, &accountv1.GetMainWalletRequest{
		AccountId: accountID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get main wallet: %w", err)
	}
	return resp.Wallet, nil
}
