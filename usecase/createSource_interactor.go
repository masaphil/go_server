package usecase

import (
	"context"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/source"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/repository"
	"github.com/quantum-box/skillforest_platform/go/services/test/usecase/boundary"
)

type CreateSourceInteractor struct {
	sourceRepo repository.SourceRepository
}

func NewCreateSourceInteractor(
	sourceRepo repository.SourceRepository,
) *CreateSourceInteractor {
	return &CreateSourceInteractor{
		sourceRepo: sourceRepo,
	}
}

func (uc *CreateSourceInteractor) Execute(ctx context.Context, sourceDto *boundary.AddSourceInputDto) error {

	quizList := make([]*source.SourceQuiz, 0, len(sourceDto.QuizList))
	for _, quiz := range sourceDto.QuizList {
		sourceQuiz, err := source.NewSourceQuiz(
			quiz.Title, quiz.Statement, source.ToOptionList(quiz.Options), source.Option(quiz.Answer))
		if err != nil {
			return err
		}
		quizList = append(quizList, sourceQuiz)
	}

	// create new source using user input data
	sourceObj := source.New(sourceDto.Title, sourceDto.Statement, quizList, sourceDto.SkillIds, sourceDto.Average, sourceDto.Hyojun)

	err := uc.sourceRepo.Create(ctx, sourceObj)
	if err != nil {
		return err
	}

	return nil
}
