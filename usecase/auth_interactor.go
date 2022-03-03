package usecase

import (
	"context"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/tester"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/repository"
)

type AuthInteractor struct {
	testerRepo repository.TesterRepository
}

func NewTesterInteractorImpl(
	testerRepo repository.TesterRepository) *AuthInteractor {
	return &AuthInteractor{
		testerRepo: testerRepo,
	}
}

//TODO
func (uc *AuthInteractor) SignUp(ctx context.Context, tester *tester.Tester) error {
	err := uc.testerRepo.Create(ctx, tester)
	return err
}
