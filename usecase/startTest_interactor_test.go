package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/paper"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/source"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/tester"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/repository"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/repository/mock_repository"
)

func TestStartTestInteractor_Execute(t *testing.T) {
	tid := tester.TesterID("")
	ctrl := gomock.NewController(t)
	mockPaperRepo := mock_repository.NewMockPaperRepository(ctrl)
	mockSourceRepo := mock_repository.NewMockSourceRepository(ctrl)
	ctx := context.Background()
	type fields struct {
		sourceRepo repository.SourceRepository
		paperRepo  repository.PaperRepository
	}
	type args struct {
		ctx context.Context
		sid source.SourceID
		tid *tester.TesterID
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *paper.Paper
		mockInit func()
		wantErr  bool
	}{
		{
			name: "test",
			fields: fields{
				sourceRepo: mockSourceRepo,
				paperRepo:  mockPaperRepo,
			},
			args: args{
				ctx: ctx,
				sid: source.SourceID(""),
				tid: &tid,
			},
			want: &paper.Paper{
				ID:       "",
				TesterId: &tid,
				SourceId: "",
				QuizList: []*paper.Quiz{{}, {}},
				Result:   nil,
			},
			mockInit: func() {
				mockSourceRepo.EXPECT().FindById(ctx, source.SourceID("")).
					Return(&source.Source{}, nil)
				mockPaperRepo.EXPECT().Create(ctx, gomock.Any()).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &StartTestInteractor{
				sourceRepo: tt.fields.sourceRepo,
				paperRepo:  tt.fields.paperRepo,
			}
			tt.mockInit()
			got, err := st.Execute(tt.args.ctx, tt.args.sid, tt.args.tid)
			if (err != nil) != tt.wantErr {
				t.Errorf("StartTestInteractor.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartTestInteractor.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
