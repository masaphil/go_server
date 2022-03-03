package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/source"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/repository"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/repository/mock_repository"
	"github.com/quantum-box/skillforest_platform/go/services/test/usecase/boundary"
)

func TestCreateSourceInteractor_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSourceRepo := mock_repository.NewMockSourceRepository(ctrl)
	ctx := context.Background()
	type fields struct {
		sourceRepo repository.SourceRepository
	}
	type args struct {
		ctx       context.Context
		sourceDto *boundary.AddSourceInputDto
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mockInit func()
		wantErr  bool
	}{
		{
			name:   "test",
			fields: fields{sourceRepo: mockSourceRepo},
			args: args{
				ctx: ctx,
				sourceDto: &boundary.AddSourceInputDto{
					Title:     "",
					Statement: "",
					QuizList: []*boundary.SourceQuizDto{
						{Statement: "問題文", Options: []string{"aaa", "bbb"}, Answer: "aaa"},
					},
					SkillIds: []string{"", ""},
					Average:  1.0,
					Hyojun:   1.0,
				},
			},
			mockInit: func() {
				mockSourceRepo.EXPECT().Create(ctx, &source.Source{
					ID:          "",
					Title:       "",
					Statement:   "",
					QuizList:    []*source.SourceQuiz{{}},
					SkillIdList: []source.SkillID{"", ""},
					Average:     1.0,
					Hyojun:      1.0,
				}).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockInit()
			uc := &CreateSourceInteractor{
				sourceRepo: tt.fields.sourceRepo,
			}
			if err := uc.Execute(tt.args.ctx, tt.args.sourceDto); (err != nil) != tt.wantErr {
				t.Errorf("AddSourceInteractor.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
