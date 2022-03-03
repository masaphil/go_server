package usecase

import (
	"context"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/source"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/repository"
)

type FindByIdSourceInteractor struct {
	sourceRepo repository.SourceRepository
}

func NewFindByIdSourceInteractor(
	sourceRepo repository.SourceRepository,
) *FindByIdSourceInteractor {
	return &FindByIdSourceInteractor{
		sourceRepo: sourceRepo,
	}
}

func (uc *FindByIdSourceInteractor) Execute(ctx context.Context, sourceId *source.SourceID) (*source.Source, error) {

	sourceObj, err := uc.sourceRepo.FindById(ctx, *sourceId)
	if err != nil {
		return nil, err
	}

	return sourceObj, nil
}
