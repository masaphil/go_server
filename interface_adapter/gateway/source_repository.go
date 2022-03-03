package gateway

import (
	"context"
	"log"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/source"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/repository"
	"github.com/quantum-box/skillforest_platform/go/services/test/interface_adapter/gateway/driver"
	"github.com/quantum-box/skillforest_platform/go/services/test/interface_adapter/gateway/rdbmodel"
)

var _ repository.SourceRepository = &sourceRepositoryImpl{}

type sourceRepositoryImpl struct {
	driver driver.PrismaDriver
}

func NewSourceRepositoryImpl(d driver.PrismaDriver) *sourceRepositoryImpl {
	/*
		defer func() {
		      if err := client.Prisma.Disconnect(); err != nil {
		          panic(err)
		      }
		  }()
	*/
	return &sourceRepositoryImpl{
		driver: d,
	}
}

func (r *sourceRepositoryImpl) Create(ctx context.Context, sce *source.Source) error {
	sourceObj, err := r.driver.DB().Source.CreateOne(
		rdbmodel.Source.ID.Set(string(sce.ID)),
		rdbmodel.Source.Title.Set(sce.Title),
		rdbmodel.Source.Stetement.Set(sce.Statement),
		rdbmodel.Source.Average.Set(float64(sce.Average)),
		rdbmodel.Source.Hyojun.Set(float64(sce.Hyojun)),
		rdbmodel.Source.Skillid.Set(source.ToStringListSkillID(sce.SkillIdList)),
	).Exec(ctx)
	if err != nil {
		return err
	}
	log.Printf("source: %+v", sourceObj)

	for _, s := range sce.QuizList {
		sourceQuizObj, err := r.driver.DB().Sourcequiz.CreateOne(
			rdbmodel.Sourcequiz.Source.Link(
				rdbmodel.Source.ID.Equals(sourceObj.ID),
			),
			rdbmodel.Sourcequiz.Title.Set(s.Title),
			rdbmodel.Sourcequiz.Statement.Set(s.Statement),
			rdbmodel.Sourcequiz.Answer.Set(string(s.Answer)),
			rdbmodel.Sourcequiz.Options.Set(source.ToStringListOption(s.Options)),
		).Exec(ctx)
		log.Printf("quiz: %+v", sourceQuizObj)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *sourceRepositoryImpl) FindAll(ctx context.Context) ([]*source.Source, error) {

	sourceRcs, err := r.driver.DB().Source.FindMany().Exec(ctx)
	if err != nil {
		return nil, err
	}

	sourceObjs := make([]*source.Source, 0, len(sourceRcs))

	for _, sourceRc := range sourceRcs {

		quizzes, err := r.driver.DB().Sourcequiz.FindMany(
			rdbmodel.Sourcequiz.SourceID.Equals(sourceRc.ID),
		).Exec(ctx)
		if err != nil {
			return nil, err
		}
		sourceQuizRcs := make([]*source.SourceQuiz, 0, len(quizzes))

		for _, sourceQuizRc := range quizzes {

			sourceQuizRcs = append(sourceQuizRcs, &source.SourceQuiz{
				Statement: sourceQuizRc.Statement,
				Options:   source.ToOptionList(sourceQuizRc.Options),
				Answer:    source.Option(sourceQuizRc.Answer),
			})

		}

		sourceObjs = append(sourceObjs, &source.Source{
			ID:          source.SourceID(sourceRc.ID),
			Title:       sourceRc.Title,
			Statement:   sourceRc.Stetement,
			QuizList:    sourceQuizRcs,
			SkillIdList: source.ToSkillIDList(sourceRc.Skillid),
			Average:     float32(sourceRc.Average),
			Hyojun:      float32(sourceRc.Hyojun),
		})
	}

	return sourceObjs, nil
}

func (r *sourceRepositoryImpl) FindById(ctx context.Context, sid source.SourceID) (*source.Source, error) {

	sourceRc, err := r.driver.DB().Source.FindUnique(
		rdbmodel.Source.ID.Equals(string(sid)),
	).With(
		rdbmodel.Source.Quizlist.Fetch(),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	quizObjs := make([]*source.SourceQuiz, 0, len(sourceRc.Quizlist()))

	for _, sourceQuizRc := range sourceRc.Quizlist() {

		quizObjs = append(quizObjs, &source.SourceQuiz{
			Title:     sourceQuizRc.Title,
			Statement: sourceQuizRc.Statement,
			Options:   source.ToOptionList(sourceQuizRc.Options),
			Answer:    source.Option(sourceQuizRc.Answer),
		})

	}

	return &source.Source{
		ID:          source.SourceID(sourceRc.ID),
		Title:       sourceRc.Title,
		Statement:   sourceRc.Stetement,
		QuizList:    quizObjs,
		SkillIdList: source.ToSkillIDList(sourceRc.Skillid),
		Average:     float32(sourceRc.Average),
		Hyojun:      float32(sourceRc.Hyojun),
	}, nil
}

func (r *sourceRepositoryImpl) Update(ctx context.Context, sourceObj *source.Source) (*source.Source, error) {

	//sourceを１件取得しアップデート
	sourceRc, err := r.driver.DB().Source.FindUnique(
		rdbmodel.Source.ID.Equals(string(sourceObj.ID)),
	).With(
		rdbmodel.Source.Quizlist.Fetch(),
	).Update(
		rdbmodel.Source.Title.Set(sourceObj.Title),
		rdbmodel.Source.Stetement.Set(sourceObj.Statement),
		rdbmodel.Source.Skillid.Set(source.ToStringListSkillID(sourceObj.SkillIdList)),
		rdbmodel.Source.Average.Set(float64(sourceObj.Average)),
		rdbmodel.Source.Hyojun.Set(float64(sourceObj.Hyojun)),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	//quizlistを削除
	for _, quizRc := range sourceRc.Quizlist() {
		_, err := r.driver.DB().Sourcequiz.FindUnique(
			rdbmodel.Sourcequiz.ID.Equals(quizRc.ID),
		).Delete().Exec(ctx)
		if err != nil {
			return nil, err
		}
	}

	//quizlistを新規作成
	for _, s := range sourceObj.QuizList {
		sourceQuizRc, err := r.driver.DB().Sourcequiz.CreateOne(
			rdbmodel.Sourcequiz.Source.Link(
				rdbmodel.Source.ID.Equals(string(sourceObj.ID)),
			),
			rdbmodel.Sourcequiz.Title.Set(s.Title),
			rdbmodel.Sourcequiz.Statement.Set(s.Statement),
			rdbmodel.Sourcequiz.Answer.Set(string(s.Answer)),
			rdbmodel.Sourcequiz.Options.Set(source.ToStringListOption(s.Options)),
		).Exec(ctx)
		log.Printf("quiz: %+v", sourceQuizRc)
		if err != nil {
			return nil, err
		}
	}

	return sourceObj, nil
}
func (r *sourceRepositoryImpl) Delete(ctx context.Context, sid source.SourceID) (*source.SourceID, error) {

	_, err := r.driver.DB().Source.FindUnique(
		rdbmodel.Source.ID.Equals(string(sid)),
	).Delete().Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &sid, nil
}
