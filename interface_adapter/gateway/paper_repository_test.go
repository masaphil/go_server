package gateway

import (
	"context"
	"testing"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/paper"
	"github.com/quantum-box/skillforest_platform/go/services/test/interface_adapter/gateway/driver"
)

func Test_paperRepositoryImpl_Create(t *testing.T) {
	type fields struct {
		driver driver.PrismaDriver
	}
	type args struct {
		ctx      context.Context
		paperObj *paper.Paper
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
			r := &paperRepositoryImpl{
				driver: tt.fields.driver,
			}
			if err := r.Create(tt.args.ctx, tt.args.paperObj); (err != nil) != tt.wantErr {
				t.Errorf("paperRepositoryImpl.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
