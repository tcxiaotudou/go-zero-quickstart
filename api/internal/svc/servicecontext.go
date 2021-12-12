package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"go-zero-quickstart/api/internal/config"
	"go-zero-quickstart/rpc/transform/transformer"
)

type ServiceContext struct {
	Config config.Config
	Transformer transformer.Transformer    // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Transformer: transformer.NewTransformer(zrpc.MustNewClient(c.Transform)), // 手动代码
	}
}
