package repository

import (
	"context"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/tester"
)

//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
type TesterRepository interface {
	Create(context.Context, *tester.Tester) error
}
