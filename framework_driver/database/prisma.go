package database

import (
	"github.com/quantum-box/skillforest_platform/go/services/test/interface_adapter/gateway/driver"
	"github.com/quantum-box/skillforest_platform/go/services/test/interface_adapter/gateway/rdbmodel"
)

var _ driver.PrismaDriver = &PrismaDriverImpl{}

type PrismaDriverImpl struct {
	client *rdbmodel.PrismaClient
}

func NewPrismaDriverImpl() *PrismaDriverImpl {
	cli := rdbmodel.NewClient()
	if err := cli.Prisma.Connect(); err != nil {
		panic(err)
	}
	println("established database connection")
	return &PrismaDriverImpl{
		client: cli,
	}
}

func (d *PrismaDriverImpl) DB() *rdbmodel.PrismaClient { return d.client }
