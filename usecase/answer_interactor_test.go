package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/paper"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/source"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/tester"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/repository"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/repository/mock_repository"
)

func TestAnswerInteractor_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPaperRepo := mock_repository.NewMockPaperRepository(ctrl)
	ctx := context.Background()
	type fields struct {
		paperRepo repository.PaperRepository
	}
	type args struct {
		ctx   context.Context
		pi    paper.PaperID
		index int
		o     paper.Option
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mockInit func()
		wantErr  bool
	}{
		{
			name:   "成功",
			fields: fields{paperRepo: mockPaperRepo},
			args: args{
				ctx:   ctx,
				pi:    paper.PaperID(""),
				index: 0,
				o:     paper.Option("aaa"),
			},
			mockInit: func() {
				tid := tester.TesterID("")
				mockPaperRepo.EXPECT().GetById(ctx, paper.PaperID("")).
					Return(&paper.Paper{
						ID:       paper.PaperID(""),
						TesterId: &tid,
						SourceId: source.SourceID(""),
						QuizList: []*paper.Quiz{{}, {}},
						Result:   &paper.Result{},
					}, nil)
				mockPaperRepo.EXPECT().Update(ctx, gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockInit()
			ai := &AnswerInteractor{
				paperRepo: tt.fields.paperRepo,
			}
			if err := ai.Execute(tt.args.ctx, tt.args.pi, tt.args.index, tt.args.o); (err != nil) != tt.wantErr {
				t.Errorf("AnswerInteractor.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
