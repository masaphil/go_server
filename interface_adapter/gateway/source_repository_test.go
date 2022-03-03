package gateway

import (
	"context"
	"testing"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/source"
	"github.com/quantum-box/skillforest_platform/go/services/test/interface_adapter/gateway/driver"
)

// func TestSourceRepositoryImpl_Save(t *testing.T) {
// 	type fields struct {
// 		client *rdbmodel.PrismaClient
// 	}
// 	type args struct {
// 		ctx context.Context
// 		sce *source.Source
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "test",
// 			fields: fields{
// 				client: rdbmodel.NewClient(),
// 			},
// 			args: args{
// 				ctx: context.Background(),
// 				sce: &source.Source{
// 					ID:        source.SourceID("123456"),
// 					Title:     "",
// 					Statement: "",
// 					QuizList: []*source.SourceQuiz{
// 						{
// 							Statement: "test問題文です。",
// 							Options:   []source.Option{"aaa", "bbb", "ccc"},
// 							Answer:    source.Option("aaa"),
// 						},
// 						{
// 							Statement: "test2問題文です。",
// 							Options:   []source.Option{"aaa", "bbb", "ccc"},
// 							Answer:    source.Option("aaa"),
// 						},
// 					},
// 					SkillIdList: []source.SkillID{"aaa", "bbb", "ccc"},
// 					Average:     1.1,
// 					Hyojun:      1.1,
// 				},
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err := tt.fields.client.Prisma.Connect()
// 			if err != nil {
// 				panic(err)
// 			}
// 			defer func() {
// 				if err := tt.fields.client.Prisma.Disconnect(); err != nil {
// 					panic(err)
// 				}
// 			}()
// 			r := &sourceRepositoryImpl{
// 				client: tt.fields.client,
// 			}
// 			if err := r.Create(tt.args.ctx, tt.args.sce); (err != nil) != tt.wantErr {
// 				t.Errorf("SourceRepositoryImpl.Save() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestSourceRepositoryImpl_FindAll(t *testing.T) {
// 	type fields struct {
// 		client *rdbmodel.PrismaClient
// 	}
// 	type args struct {
// 		ctx context.Context
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "test",
// 			fields: fields{
// 				client: rdbmodel.NewClient(),
// 			},
// 			args: args{
// 				ctx: context.Background(),
// 			},
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err := tt.fields.client.Prisma.Connect()
// 			if err != nil {
// 				panic(err)
// 			}
// 			defer func() {
// 				if err := tt.fields.client.Prisma.Disconnect(); err != nil {
// 					panic(err)
// 				}
// 			}()
// 			r := &sourceRepositoryImpl{
// 				client: tt.fields.client,
// 			}
// 			if rec, err := r.FindAll(tt.args.ctx); (err != nil) != tt.wantErr {
// 				mar, _ := json.Marshal(*rec[0])
// 				fmt.Printf(string(mar))
// 				t.Errorf("TestSourceRepositoryImpl_FindAll() error = %v, wantErr %v", *rec[0].QuizList[0], tt.wantErr)
// 			}
// 		})
// 	}
// }

// func Test_sourceRepositoryImpl_FindById(t *testing.T) {
// 	type fields struct {
// 		client *rdbmodel.PrismaClient
// 	}
// 	type args struct {
// 		ctx context.Context
// 		sid source.SourceID
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *source.Source
// 		wantErr bool
// 	}{
// 		{
// 			name: "test",
// 			fields: fields{
// 				client: rdbmodel.NewClient(),
// 			},
// 			args: args{
// 				ctx: context.Background(),
// 				sid: source.SourceID("1232345"),
// 			},
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err := tt.fields.client.Prisma.Connect()
// 			if err != nil {
// 				panic(err)
// 			}
// 			defer func() {
// 				if err := tt.fields.client.Prisma.Disconnect(); err != nil {
// 					panic(err)
// 				}
// 			}()
// 			r := &sourceRepositoryImpl{
// 				client: tt.fields.client,
// 			}
// 			got, err := r.FindById(tt.args.ctx, tt.args.sid)

// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("sourceRepositoryImpl.FindById() error = %v, wantErr %v", got.QuizList[0].Title, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("sourceRepositoryImpl.FindById() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_sourceRepositoryImpl_Update(t *testing.T) {
// 	type fields struct {
// 		client *rdbmodel.PrismaClient
// 	}
// 	type args struct {
// 		ctx context.Context
// 		sce *source.Source
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := &sourceRepositoryImpl{
// 				client: tt.fields.client,
// 			}
// 			if _, err := r.Update(tt.args.ctx, tt.args.sce); (err != nil) != tt.wantErr {
// 				t.Errorf("sourceRepositoryImpl.Update() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func Test_sourceRepositoryImpl_Delete(t *testing.T) {
// 	type fields struct {
// 		client *rdbmodel.PrismaClient
// 	}
// 	type args struct {
// 		ctx context.Context
// 		sid source.SourceID
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *source.SourceID
// 		wantErr bool
// 	}{
// 		{
// 			name: "test",
// 			fields: fields{
// 				client: rdbmodel.NewClient(),
// 			},
// 			args: args{
// 				ctx: context.Background(),
// 				sid: source.SourceID("1232347"),
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err := tt.fields.client.Prisma.Connect()
// 			if err != nil {
// 				panic(err)
// 			}
// 			defer func() {
// 				if err := tt.fields.client.Prisma.Disconnect(); err != nil {
// 					panic(err)
// 				}
// 			}()
// 			r := &sourceRepositoryImpl{
// 				client: tt.fields.client,
// 			}
// 			got, err := r.Delete(tt.args.ctx, tt.args.sid)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("sourceRepositoryImpl.Delete() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("sourceRepositoryImpl.Delete() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
func Test_sourceRepositoryImpl_Create(t *testing.T) {
	type fields struct {
		driver driver.PrismaDriver
	}
	type args struct {
		ctx context.Context
		sce *source.Source
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &sourceRepositoryImpl{
				driver: tt.fields.driver,
			}
			if err := r.Create(tt.args.ctx, tt.args.sce); (err != nil) != tt.wantErr {
				t.Errorf("sourceRepositoryImpl.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
