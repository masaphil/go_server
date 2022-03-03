package gateway

import (
	"context"
	"log"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/paper"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/source"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/tester"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/repository"
	"github.com/quantum-box/skillforest_platform/go/services/test/interface_adapter/gateway/driver"
	"github.com/quantum-box/skillforest_platform/go/services/test/interface_adapter/gateway/rdbmodel"
)

var _ repository.PaperRepository = &paperRepositoryImpl{}

type paperRepositoryImpl struct {
	driver driver.PrismaDriver
}

func NewPaperRepositoryImpl(d driver.PrismaDriver) *paperRepositoryImpl {
	return &paperRepositoryImpl{
		driver: d,
	}
}

func (r *paperRepositoryImpl) Create(ctx context.Context, paperObj *paper.Paper) error {
	paperRc, err := r.driver.DB().Paper.CreateOne(
		rdbmodel.Paper.ID.Set(string(paperObj.ID)),
		rdbmodel.Paper.SourceID.Set(string(paperObj.SourceId)),
		rdbmodel.Paper.TesterID.Set(string(*paperObj.TesterId)),
	).Exec(ctx)
	log.Printf("quiz: %+v", paperRc)
	if err != nil {
		return err
	}

	resultRc, err := r.driver.DB().Result.CreateOne(
		rdbmodel.Result.Paper.Link(
			rdbmodel.Paper.ID.Equals(paperRc.ID),
		),
		rdbmodel.Result.Volume.Set(paperObj.Result.Volume),
		rdbmodel.Result.Average.Set(float64(paperObj.Result.Volume)),
		rdbmodel.Result.Hyojun.Set(float64(paperObj.Result.Hyojun)),
		rdbmodel.Result.Point.Set(*paperObj.Result.Point),
		rdbmodel.Result.Deviation.Set(float64(*paperObj.Result.Deviation)),
	).Exec(ctx)
	log.Printf("quiz: %+v", resultRc)
	if err != nil {
		return err
	}

	for _, quiz := range paperObj.QuizList {
		paperQuizRc, err := r.driver.DB().PaperQuiz.CreateOne(
			rdbmodel.PaperQuiz.Paper.Link(
				rdbmodel.Paper.ID.Equals(paperRc.ID),
			),
			rdbmodel.PaperQuiz.Statement.Set(quiz.Statement),
			rdbmodel.PaperQuiz.Answer.Set(string(quiz.Answer)),
			rdbmodel.PaperQuiz.Answered.Set(string(*quiz.Answered)),
			rdbmodel.PaperQuiz.IsReadAnswer.Set(quiz.IsReadAnswer),
		).Exec(ctx)
		log.Printf("quiz: %+v", paperQuizRc)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *paperRepositoryImpl) GetById(ctx context.Context, paperId *paper.PaperID) (*paper.Paper, error) {
	paperRc, err := r.driver.DB().Paper.FindUnique(
		rdbmodel.Paper.ID.Equals(string(*paperId)),
	).With(
		rdbmodel.Paper.Result.Fetch(),
		rdbmodel.Paper.QuizList.Fetch(),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	quizObjs := make([]*paper.Quiz, 0, len(paperRc.QuizList()))
	for _, quizRc := range paperRc.QuizList() {
		op := paper.Option(quizRc.Answered)
		quizObjs = append(quizObjs, &paper.Quiz{
			Statement:    quizRc.Statement,
			Options:      paper.ToOptionList(quizRc.Options),
			Answer:       paper.Option(quizRc.Answer),
			Answered:     &op,
			IsReadAnswer: quizRc.IsReadAnswer,
		})
	}
	resultRc, _ := paperRc.Result()
	pointRc, _ := resultRc.Point()
	dev, _ := resultRc.Deviation()
	deviationRc := float32(dev)
	resultObj := &paper.Result{
		Point:     &pointRc,
		Deviation: &deviationRc,
		Volume:    resultRc.Volume,
		Average:   float32(resultRc.Average),
		Hyojun:    float32(resultRc.Hyojun),
	}

	tidf, _ := paperRc.TesterID()
	tid := tester.TesterID(tidf)
	return paper.New(
		paper.PaperID(paperRc.ID),
		&tid,
		source.SourceID(paperRc.SourceID),
		quizObjs,
		resultObj,
	), nil
}

func (r *paperRepositoryImpl) Update(ctx context.Context, paperObj *paper.Paper) (*paper.Paper, error) {
	paperRc, err := r.driver.DB().Paper.FindUnique(
		rdbmodel.Paper.ID.Equals(string(paperObj.ID)),
	).With(
		rdbmodel.Paper.Result.Fetch(),
		rdbmodel.Paper.QuizList.Fetch(),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	//delete quizlist
	for _, quizRc := range paperRc.QuizList() {

		_, err := r.driver.DB().PaperQuiz.FindUnique(
			rdbmodel.PaperQuiz.ID.Equals(quizRc.ID),
		).Delete().Exec(ctx)
		if err != nil {
			return nil, err
		}
	}
	//delete result
	_, err = r.driver.DB().Result.FindUnique(
		rdbmodel.Result.ResultID.Equals(string(paperObj.ID)),
	).Delete().Exec(ctx)
	if err != nil {
		return nil, err
	}

	//paperQuizインサート
	for _, quiz := range paperObj.QuizList {
		quizRc, err := r.driver.DB().PaperQuiz.CreateOne(
			rdbmodel.PaperQuiz.Paper.Link(
				rdbmodel.Paper.ID.Equals(string(paperObj.ID)),
			),
			rdbmodel.PaperQuiz.Statement.Set(quiz.Statement),
			rdbmodel.PaperQuiz.Answer.Set(string(quiz.Answer)),
			rdbmodel.PaperQuiz.Answered.Set(string(*quiz.Answered)),
			rdbmodel.PaperQuiz.IsReadAnswer.Set(quiz.IsReadAnswer),
		).Exec(ctx)
		if err != nil {
			return nil, err
		}
		log.Printf("paperquiz: %+v", quizRc)
	}

	//Resultインサート
	resultRc, err := r.driver.DB().Result.CreateOne(
		rdbmodel.Result.Paper.Link(
			rdbmodel.Paper.ID.Equals(string(paperObj.ID)),
		),
		rdbmodel.Result.Volume.Set(paperObj.Result.Volume),
		rdbmodel.Result.Average.Set(float64(paperObj.Result.Volume)),
		rdbmodel.Result.Hyojun.Set(float64(paperObj.Result.Hyojun)),
		rdbmodel.Result.Point.Set(*paperObj.Result.Point),
		rdbmodel.Result.Deviation.Set(float64(*paperObj.Result.Deviation)),
	).Exec(ctx)
	log.Printf("result: %+v", resultRc)
	if err != nil {
		return nil, err
	}

	return paperObj, nil
}
