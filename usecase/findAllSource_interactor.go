package usecase

import (
	"context"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/source"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/repository"
)

type FindAllSourceInteractor struct {
	sourceRepo repository.SourceRepository
}

func NewfindAllSourceInteractor(
	sourceRepo repository.SourceRepository,
) *FindAllSourceInteractor {
	return &FindAllSourceInteractor{
		sourceRepo: sourceRepo,
	}
}

func (uc *FindAllSourceInteractor) Execute(ctx context.Context) ([]*source.Source, error) {

	sourceList, err := uc.sourceRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return sourceList, nil
}
