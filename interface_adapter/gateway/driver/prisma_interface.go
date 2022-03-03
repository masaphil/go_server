package driver

import "github.com/quantum-box/skillforest_platform/go/services/test/interface_adapter/gateway/rdbmodel"

type PrismaDriver interface {
	DB() *rdbmodel.PrismaClient
}
