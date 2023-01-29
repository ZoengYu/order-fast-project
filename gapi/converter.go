package gapi

import (
	db "github.com/ZoengYu/order-fast-project/db/sqlc"
	"github.com/ZoengYu/order-fast-project/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertDBUserToPb(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
