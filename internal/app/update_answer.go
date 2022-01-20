package app

import "context"

func (m *MicroserviceServer) UpdateAnswer(ctx context.Context, req *desc.UpdateAnswerRequest) (*desc.UpdateAnswerResponse, error) {
	userID, err := m.getUserIdFromToken(ctx)
	if err !=nil{
		return nil, err 
	}

	answer, err := m.answerService.UpdateAnswer(dto.Answer{
		UserId:		userId,
		ID:			req.GetId(),
		Answer:		req.GetAnswer(),
		Score:		req.GetScore(),
	})
	if err !=nil {
		return nil, err 
	}
	return &desc.UpdateAnswerResponse{
		Id:			answer.ID,
		QuestionID:	answer.QuestionID,
		Answer:		answer.Answer,
		Score:		req.GetScore()}, nil
	}
}