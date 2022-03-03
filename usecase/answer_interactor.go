package usecase

import (
	"context"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/paper"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/repository"
)

type AnswerInteractor struct {
	paperRepo repository.PaperRepository
}

func NewAnswerInteractor(
	paperRepo repository.PaperRepository,
) *AnswerInteractor {
	return &AnswerInteractor{
		paperRepo: paperRepo,
	}
}

//paperの中の指定のpaperQuizの回答をする
func (uc *AnswerInteractor) Execute(ctx context.Context, pi paper.PaperID, index int, o paper.Option) error {

	paperObj, err := uc.paperRepo.GetById(ctx, &pi)
	if err != nil {
		return err
	}

	err = paperObj.Answer(index, o)
	if err != nil {
		return err
	}

	_, err = uc.paperRepo.Update(ctx, paperObj)
	if err != nil {
		return err
	}

	return nil
}
