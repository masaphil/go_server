package usecase

import (
	"context"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/source"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/repository"
)

type UpdateSourceInteractor struct {
	sourceRepo repository.SourceRepository
}

func NewUpdateSourceInteractor(
	sourceRepo repository.SourceRepository,
) *UpdateSourceInteractor {
	return &UpdateSourceInteractor{
		sourceRepo: sourceRepo,
	}
}

func (uc *UpdateSourceInteractor) Execute(ctx context.Context, source *source.Source) (*source.Source, error) {

	sourceRc, err := uc.sourceRepo.Update(ctx, source)
	if err != nil {
		return nil, err
	}

	return sourceRc, nil
}
