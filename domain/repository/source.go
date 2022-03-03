package repository

import (
	"context"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/source"
)

//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

type SourceRepository interface {
	Create(context.Context, *source.Source) error
	FindAll(context.Context) ([]*source.Source, error)
	FindById(context.Context, source.SourceID) (*source.Source, error)
	Update(context.Context, *source.Source) (*source.Source, error)
	Delete(context.Context, source.SourceID) (*source.SourceID, error)
}
