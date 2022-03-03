package controller

import (
	"context"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/paper"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/source"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/tester"
	"github.com/quantum-box/skillforest_platform/go/services/test/interface_adapter/controller/pb"
	"github.com/quantum-box/skillforest_platform/go/services/test/usecase"
)

type PaperHandlerImpl struct {
	pb.UnimplementedPaperApiServer

	startTestUse     *usecase.StartTestInteractor
	findByIdPaperUse *usecase.FindByIdPaperInteractor
	updatePaperUse   *usecase.UpdatePaperInterctor
}

var _ pb.PaperApiServer = &PaperHandlerImpl{}

func NewPaperHandler(
	startTest *usecase.StartTestInteractor,
	findById *usecase.FindByIdPaperInteractor,
	update *usecase.UpdatePaperInterctor,
) *PaperHandlerImpl {
	return &PaperHandlerImpl{
		startTestUse:     startTest,
		findByIdPaperUse: findById,
		updatePaperUse:   update,
	}
}
func convertToResponsePaper(paperObj *paper.Paper) *pb.Paper {

	resquizzes := make([]*pb.PaperQuiz, 0, len(paperObj.QuizList))
	for _, quiz := range paperObj.QuizList {
		resquizzes = append(resquizzes, &pb.PaperQuiz{
			Statement:    quiz.Statement,
			Options:      paper.ToStringListOption(quiz.Options),
			Answer:       string(quiz.Answer),
			Answered:     string(*quiz.Answered),
			IsReadAnswer: quiz.IsReadAnswer,
		})
	}
	resResult := &pb.Result{
		Point:     int32(*paperObj.Result.Point),
		Deviation: *paperObj.Result.Deviation,
		Volume:    int32(paperObj.Result.Volume),
		Average:   paperObj.Result.Average,
		Hyojun:    paperObj.Result.Hyojun,
	}

	testerId := string(*paperObj.TesterId)
	return &pb.Paper{
		Id:           string(paperObj.ID),
		TesterId:     testerId,
		SourceId:     string(paperObj.SourceId),
		PaperQuizzes: resquizzes,
		Result:       resResult,
	}
}
func convertToPaperObj(req *pb.Paper) *paper.Paper {

	quizObjs := make([]*paper.Quiz, 0, len(req.PaperQuizzes))
	for _, quiz := range req.PaperQuizzes {
		quizObjs = append(quizObjs, &paper.Quiz{
			Statement:    quiz.Statement,
			Options:      paper.ToOptionList(quiz.Options),
			Answer:       paper.Option(quiz.Answer),
			Answered:     (*paper.Option)(&quiz.Answered),
			IsReadAnswer: quiz.IsReadAnswer,
		})
	}
	po := int(req.Result.Point)
	resultObj := paper.Result{
		Point:     &po,
		Deviation: &req.Result.Deviation,
		Volume:    int(req.Result.Volume),
		Average:   req.Result.Average,
		Hyojun:    req.Result.Hyojun,
	}
	return &paper.Paper{
		ID:       paper.PaperID(req.Id),
		TesterId: (*tester.TesterID)(&req.TesterId),
		SourceId: source.SourceID(req.SourceId),
		QuizList: quizObjs,
		Result:   &resultObj,
	}
}

func (h *PaperHandlerImpl) Start(ctx context.Context, req *pb.StartTestRequest) (*pb.StartTestResponse, error) {

	sid := source.SourceID(req.SourceId)
	tid := tester.TesterID(*req.UserId)
	paperObj, err := h.startTestUse.Execute(ctx, sid, &tid)
	if err != nil {
		return nil, err
	}

	return &pb.StartTestResponse{
		Paper: convertToResponsePaper(paperObj),
	}, nil

}

func (h *PaperHandlerImpl) FindById(ctx context.Context, req *pb.FindByIdPaperRequest) (*pb.FindByIdPaperResponse, error) {

	pid := paper.PaperID(req.PaperId)
	paperObj, err := h.findByIdPaperUse.Execute(ctx, &pid)
	if err != nil {
		return nil, err
	}

	return &pb.FindByIdPaperResponse{
		Paper: convertToResponsePaper(paperObj),
	}, nil

}

func (h *PaperHandlerImpl) Update(ctx context.Context, req *pb.UpdatePaperRequest) (*pb.UpdatePaperResponse, error) {

	paperObj, err := h.updatePaperUse.Execute(ctx, convertToPaperObj(req.Paper))
	if err != nil {
		return nil, err
	}

	return &pb.UpdatePaperResponse{
		Paper: convertToResponsePaper(paperObj),
	}, nil

}
