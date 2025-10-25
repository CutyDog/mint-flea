package server

import (
	"context"

	accountv1 "github.com/CutyDog/mint-flea/proto/gen/account/v1"
	"github.com/CutyDog/mint-flea/services/account/internal/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccountServer struct {
	accountv1.UnimplementedAccountServiceServer
	repo repo.AccountRepository
}

func NewAccountServer(repository repo.AccountRepository) *AccountServer {
	return &AccountServer{repo: repository}
}

func (s *AccountServer) LoginAccount(ctx context.Context, req *accountv1.LoginAccountRequest) (*accountv1.LoginAccountResponse, error) {
	uid := req.GetUid()
	if uid == "" {
		return nil, status.Error(codes.InvalidArgument, "uid must be provided")
	}

	acc, err := s.repo.FindByUID(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get account: %v", err)
	}
	if acc == nil {
		// アカウントが存在しない場合は新規作成
		acc, err = s.repo.Create(ctx, uid)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to create account: %v", err)
		}
	}

	account := &accountv1.Account{
		Id:        int64(acc.ID),
		Uid:       acc.UID,
		CreatedAt: timestamppb.New(acc.CreatedAt),
		UpdatedAt: timestamppb.New(acc.UpdatedAt),
	}

	return &accountv1.LoginAccountResponse{
		Account: account,
	}, nil
}
