package app

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	desc "github.com/yasir2000/my-do-dev-side-project-3/pkg"
)

func (m *MicroserviceServer) DeleteFile(ctx context.Context, req *desc.DeleteFileRequest) (*emptypb.Empty, error) {
	err:= m.fileUploaderService.DeleteFile(req.GetFilePath())
	if err := nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}