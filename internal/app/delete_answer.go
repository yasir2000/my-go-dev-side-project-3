package app

import (
	"context"

	desc "github.com/yasir2000/my-do-dev-side-project-3/pkg"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (m *MicroserviceServer) DeleteAnswer(ctx context.Context, req *desc.DeleteAswerRequest) (*emptypb.Empty, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	err := m.answerService.DeleteAnswer(userID, req.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
