package usecase

import (
	"context"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/paper"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/repository"
)

//paperが試験を終了する(試験時間、最後の問題をとき終えたなど)
type FinishTestInteractor struct {
	paperRepo repository.PaperRepository
}

func NewFinishTestInteractor(
	paperRepo repository.PaperRepository,
) *FinishTestInteractor {
	return &FinishTestInteractor{
		paperRepo: paperRepo,
	}
}

func (fi *FinishTestInteractor) Execute(ctx context.Context, pi paper.PaperID) error {

	paperObj, err := fi.paperRepo.GetById(ctx, &pi)
	if err != nil {
		return err
	}

	paperObj = paperObj.FinishTest()

	_, err = fi.paperRepo.Update(ctx, paperObj)
	if err != nil {
		return err
	}

	return nil
}
