package repository

//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

import (
	"context"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/paper"
)

type PaperRepository interface {
	Create(context.Context, *paper.Paper) error
	GetById(context.Context, *paper.PaperID) (*paper.Paper, error)
	Update(context.Context, *paper.Paper) (*paper.Paper, error)
}
