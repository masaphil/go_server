package usecase

import (
	"context"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/paper"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/repository"
)

type FindByIdPaperInteractor struct {
	paperRepo repository.PaperRepository
}

func NewFindByIdPaperInteractor(
	paperRepo repository.PaperRepository,
) *FindByIdPaperInteractor {
	return &FindByIdPaperInteractor{
		paperRepo: paperRepo,
	}
}

func (uc *FindByIdPaperInteractor) Execute(ctx context.Context, pid *paper.PaperID) (*paper.Paper, error) {

	paperObj, err := uc.paperRepo.GetById(ctx, pid)
	if err != nil {
		return nil, err
	}

	return paperObj, nil
}
