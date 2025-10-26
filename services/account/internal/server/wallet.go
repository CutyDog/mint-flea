package server

import (
	"context"

	accountv1 "github.com/CutyDog/mint-flea/proto/gen/account/v1"
	"github.com/CutyDog/mint-flea/services/account/internal/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WalletServer struct {
	accountv1.UnimplementedWalletServiceServer
	repo repo.WalletRepository
}

func NewWalletServer(repository repo.WalletRepository) *WalletServer {
	return &WalletServer{repo: repository}
}

func (s *WalletServer) ListWallets(ctx context.Context, req *accountv1.ListWalletsRequest) (*accountv1.ListWalletsResponse, error) {
	wallets, err := s.repo.ListByAccountID(ctx, req.GetAccountId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list wallets: %v", err)
	}

	walletsPb := make([]*accountv1.Wallet, len(wallets))
	for i, wallet := range wallets {
		walletsPb[i] = &accountv1.Wallet{
			Id:        int64(wallet.ID),
			AccountId: wallet.AccountID,
			Address:   wallet.Address,
			ChainId:   wallet.ChainID,
			IsMain:    wallet.IsMain,
			CreatedAt: timestamppb.New(wallet.CreatedAt),
			UpdatedAt: timestamppb.New(wallet.UpdatedAt),
		}
	}

	return &accountv1.ListWalletsResponse{
		Wallets: walletsPb,
	}, nil
}

func (s *WalletServer) LinkWallet(ctx context.Context, req *accountv1.LinkWalletRequest) (*accountv1.LinkWalletResponse, error) {
	wallet, err := s.repo.Create(ctx, repo.CreateWalletParams{
		AccountID: req.GetAccountId(),
		Address:   req.GetAddress(),
		ChainID:   req.GetChainId(),
		IsMain:    req.GetIsMain(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to link wallet: %v", err)
	}

	return &accountv1.LinkWalletResponse{
		Wallet: &accountv1.Wallet{
			Id:        int64(wallet.ID),
			AccountId: wallet.AccountID,
			Address:   wallet.Address,
			ChainId:   wallet.ChainID,
			IsMain:    wallet.IsMain,
			CreatedAt: timestamppb.New(wallet.CreatedAt),
			UpdatedAt: timestamppb.New(wallet.UpdatedAt),
		},
	}, nil
}

func (s *WalletServer) UnlinkWallet(ctx context.Context, req *accountv1.UnlinkWalletRequest) (*accountv1.UnlinkWalletResponse, error) {
	success, err := s.repo.Delete(ctx, req.GetWalletId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unlink wallet: %v", err)
	}

	return &accountv1.UnlinkWalletResponse{
		Success: success,
	}, nil
}

func (s *WalletServer) SetMainWallet(ctx context.Context, req *accountv1.SetMainWalletRequest) (*accountv1.SetMainWalletResponse, error) {
	wallet, err := s.repo.SetMain(ctx, req.GetWalletId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to set main wallet: %v", err)
	}

	return &accountv1.SetMainWalletResponse{
		Wallet: &accountv1.Wallet{
			Id:        int64(wallet.ID),
			AccountId: wallet.AccountID,
			Address:   wallet.Address,
			ChainId:   wallet.ChainID,
			IsMain:    wallet.IsMain,
			CreatedAt: timestamppb.New(wallet.CreatedAt),
			UpdatedAt: timestamppb.New(wallet.UpdatedAt),
		},
	}, nil
}

func (s *WalletServer) GetMainWallet(ctx context.Context, req *accountv1.GetMainWalletRequest) (*accountv1.GetMainWalletResponse, error) {
	wallet, err := s.repo.GetMain(ctx, req.GetAccountId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get main wallet: %v", err)
	}

	return &accountv1.GetMainWalletResponse{
		Wallet: &accountv1.Wallet{
			Id:        int64(wallet.ID),
			AccountId: wallet.AccountID,
			Address:   wallet.Address,
			ChainId:   wallet.ChainID,
			IsMain:    wallet.IsMain,
			CreatedAt: timestamppb.New(wallet.CreatedAt),
			UpdatedAt: timestamppb.New(wallet.UpdatedAt),
		},
	}, nil
}
