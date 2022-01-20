package app

import (
	"context"

	"github.com/yasir2000/my-do-dev-side-project-3/internal/dto"
	desc "github.com/yasir2000/my-do-dev-side-project-3/pkg"
)

func (m *MicroserviceServer) CreateAnswer(ctx context.Context, req *desc.CreateAnswerRequest) (*desc.CreateAnswerResponse, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err != nil {
		return nil, err
	}

	id, err := m.answerService.CreateAnswer(dto.Answer{
		UserID:     userID,
		QuestionID: req.GetQuestionId(),
		Answer:     req.GetAnswer(),
		Score:      req.GetScore(),
	})
	if err != nil {
		return nil, err
	}
	return &desc.CreateAnswerResponse{Id: *id}, nil
}
