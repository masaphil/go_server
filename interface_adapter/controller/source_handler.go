package controller

import (
	"context"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/source"
	"github.com/quantum-box/skillforest_platform/go/services/test/interface_adapter/controller/pb"
	"github.com/quantum-box/skillforest_platform/go/services/test/usecase"
	"github.com/quantum-box/skillforest_platform/go/services/test/usecase/boundary"
)

type SourceHandlerImpl struct {
	pb.UnimplementedSourceApiServer

	createSourceUse   *usecase.CreateSourceInteractor
	findAllSourceUse  *usecase.FindAllSourceInteractor
	findByIdSourceUse *usecase.FindByIdSourceInteractor
	updateSourceUse   *usecase.UpdateSourceInteractor
	deleteSourceUse   *usecase.DeleteSourceInteractor
}

var _ pb.SourceApiServer = &SourceHandlerImpl{}

func NewSourceHandler(
	addSource *usecase.CreateSourceInteractor,
	findAll *usecase.FindAllSourceInteractor,
	findById *usecase.FindByIdSourceInteractor,
	update *usecase.UpdateSourceInteractor,
	delete *usecase.DeleteSourceInteractor,
) *SourceHandlerImpl {
	return &SourceHandlerImpl{
		createSourceUse:   addSource,
		findAllSourceUse:  findAll,
		findByIdSourceUse: findById,
		updateSourceUse:   update,
		deleteSourceUse:   delete,
	}
}

func (h *SourceHandlerImpl) Create(ctx context.Context, req *pb.AddSourceRequest) (*pb.AddSourceResponse, error) {

	quizDtoList := make([]*boundary.SourceQuizDto, 0, len(req.Source.QuizList))

	for _, s := range req.Source.QuizList {
		quizDto := boundary.SourceQuizDto{
			Statement: s.Statement,
			Options:   s.Options,
			Answer:    s.Answer,
		}
		quizDtoList = append(quizDtoList, &quizDto)
	}

	err := h.createSourceUse.Execute(ctx, &boundary.AddSourceInputDto{
		Title:     req.Source.Title,
		Statement: req.Source.Statement,
		QuizList:  quizDtoList,
		SkillIds:  req.Source.Skillids,
		Average:   req.Source.Average,
		Hyojun:    req.Source.Hyojun,
	})
	if err != nil {
		return nil, err
	}

	return &pb.AddSourceResponse{
		Source: req.Source,
	}, nil
}

func (h *SourceHandlerImpl) FindAll(ctx context.Context, empty *pb.Empty) (*pb.FindAllSourceResponse, error) {

	sourceObjs, err := h.findAllSourceUse.Execute(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]*pb.Source, 0, len(sourceObjs))
	for _, sourceObj := range sourceObjs {

		quizList := make([]*pb.SourceQuiz, 0, len(sourceObj.QuizList))
		for _, quiz := range sourceObj.QuizList {

			quizList = append(quizList, &pb.SourceQuiz{
				Title:     quiz.Title,
				Statement: quiz.Statement,
				Options:   source.ToStringListOption(quiz.Options),
				Answer:    string(quiz.Answer),
			})
		}

		res = append(res, &pb.Source{
			Id:        string(sourceObj.ID),
			Title:     sourceObj.Title,
			Statement: sourceObj.Statement,
			QuizList:  quizList,
			Skillids:  source.ToStringListSkillID(sourceObj.SkillIdList),
			Average:   sourceObj.Average,
			Hyojun:    sourceObj.Hyojun,
		})
	}

	return &pb.FindAllSourceResponse{
		Source: res,
	}, nil
}

func (h *SourceHandlerImpl) FindById(ctx context.Context, req *pb.FindByIdSourceRequest) (*pb.FindByIdSourceResponse, error) {

	sid := source.SourceID(req.Id)
	sourceObj, err := h.findByIdSourceUse.Execute(ctx, &sid)
	if err != nil {
		return nil, err
	}

	quizList := make([]*pb.SourceQuiz, 0, len(sourceObj.QuizList))
	for _, quiz := range sourceObj.QuizList {

		quizList = append(quizList, &pb.SourceQuiz{
			Title:     quiz.Title,
			Statement: quiz.Statement,
			Options:   source.ToStringListOption(quiz.Options),
			Answer:    string(quiz.Answer),
		})
	}

	return &pb.FindByIdSourceResponse{
		Source: &pb.Source{
			Id:        string(sourceObj.ID),
			Title:     sourceObj.Title,
			Statement: sourceObj.Statement,
			QuizList:  quizList,
			Skillids:  source.ToStringListSkillID(sourceObj.SkillIdList),
			Average:   sourceObj.Average,
			Hyojun:    sourceObj.Hyojun,
		},
	}, nil
}

func (h *SourceHandlerImpl) Update(ctx context.Context, req *pb.UpdateSourceRequest) (*pb.UpdateSourceResponse, error) {

	quizList := make([]*source.SourceQuiz, 0, len(req.Source.QuizList))

	for _, s := range req.Source.QuizList {
		quizObj, err := source.NewSourceQuiz(
			s.Title,
			s.Statement,
			source.ToOptionList(s.Options),
			source.Option(s.Answer),
		)
		if err != nil {
			return nil, err
		}
		quizList = append(quizList, quizObj)
	}

	_, err := h.updateSourceUse.Execute(ctx, source.NewSource(
		req.Source.Id,
		req.Source.Title,
		req.Source.Statement,
		quizList,
		req.Source.Skillids,
		req.Source.Average,
		req.Source.Hyojun,
	))
	if err != nil {
		return nil, err
	}

	return &pb.UpdateSourceResponse{
		Source: req.Source,
	}, nil
}

func (h *SourceHandlerImpl) Delete(ctx context.Context, req *pb.DeleteSourceRequest) (*pb.DeleteSourceResponse, error) {

	sid := source.SourceID(req.Id)
	_, err := h.deleteSourceUse.Execute(ctx, &sid)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteSourceResponse{
		Id: req.Id,
	}, nil
}
