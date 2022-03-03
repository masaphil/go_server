package usecase

import (
	"context"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/source"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/repository"
)

type DeleteSourceInteractor struct {
	sourceRepo repository.SourceRepository
}

func NewDeleteSourceInteractor(
	sourceRepo repository.SourceRepository,
) *DeleteSourceInteractor {
	return &DeleteSourceInteractor{
		sourceRepo: sourceRepo,
	}
}

func (uc *DeleteSourceInteractor) Execute(ctx context.Context, sid *source.SourceID) (*source.SourceID, error) {

	idRc, err := uc.sourceRepo.Delete(ctx, *sid)
	if err != nil {
		return nil, err
	}

	return idRc, nil
}
