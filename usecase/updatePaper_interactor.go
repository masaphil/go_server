package usecase

import (
	"context"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/paper"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/repository"
)

type UpdatePaperInterctor struct {
	paperRepo repository.PaperRepository
}

func NewUpdatePaperInteractor(
	paperRepo repository.PaperRepository,
) *UpdatePaperInterctor {
	return &UpdatePaperInterctor{
		paperRepo: paperRepo,
	}
}

func (uc *UpdatePaperInterctor) Execute(ctx context.Context, paperObj *paper.Paper) (*paper.Paper, error) {

	paperRc, err := uc.paperRepo.Update(ctx, paperObj)
	if err != nil {
		return nil, err
	}
	return paperRc, nil
}
