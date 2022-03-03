package usecase

import (
	"context"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/paper"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/repository"
)

type ShowAnswersInteractor struct {
	paperRepo repository.PaperRepository
}

func NewShowAnswerInteractor(
	paperRepo repository.PaperRepository,
) *ShowAnswersInteractor {
	return &ShowAnswersInteractor{
		paperRepo: paperRepo,
	}
}

//引数をtesterIDにする？？？
func (uc *ShowAnswersInteractor) Execute(ctx context.Context, pi paper.PaperID) (*paper.Paper, error) {

	paperObj, err := uc.paperRepo.GetById(ctx, &pi)
	if err != nil {
		return nil, err
	}

	return paperObj, nil
}
