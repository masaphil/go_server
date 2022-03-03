package usecase

import (
	"context"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/paper"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/source"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/tester"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/repository"
)

type StartTestInteractor struct {
	sourceRepo repository.SourceRepository
	paperRepo  repository.PaperRepository
}

func NewStartTestInteractor(
	sourceRepo repository.SourceRepository,
	paperRepo repository.PaperRepository) *StartTestInteractor {
	return &StartTestInteractor{
		sourceRepo: sourceRepo,
		paperRepo:  paperRepo,
	}
}

//sourceからpaperを作成
func (uc *StartTestInteractor) Execute(ctx context.Context, sid source.SourceID, tid *tester.TesterID) (*paper.Paper, error) {

	sourceObj, err := uc.sourceRepo.FindById(ctx, sid)
	if err != nil {
		return nil, err
	}

	paperObj := paper.NewPaperFromSource(tid, sourceObj)

	err = uc.paperRepo.Create(ctx, paperObj)
	if err != nil {
		return nil, err
	}

	return paperObj, nil

}
