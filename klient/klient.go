package klient

import (
	"context"
	"github.com/zbitech/common/pkg/model/ztypes"
	"time"

	"github.com/zbitech/common/interfaces"
	"github.com/zbitech/common/pkg/logger"
	"github.com/zbitech/common/pkg/rctx"
	klient "github.com/zbitech/fake/klient/k8s-client"
	zklient "github.com/zbitech/fake/klient/zbi-klient"
)

type KlientFactory struct {
	zbiKlient interfaces.ZBIClientIF
	klient    interfaces.KlientIF
}

type FakeKlientFactory struct {
	zbiKlient interfaces.ZBIClientIF
	klient    interfaces.KlientIF
}

func NewFakeKlientFactory() interfaces.KlientFactoryIF {
	return &FakeKlientFactory{}
}

func (k *FakeKlientFactory) Init(ctx context.Context, namespace string, rtypes ...ztypes.ResourceObjectType) error {
	ctx = rctx.BuildContext(ctx, rctx.Context(rctx.Component, "KlientFactory"), rctx.Context(rctx.StartTime, time.Now()))
	defer logger.LogComponentTime(ctx)

	var err error

	klient, err := klient.NewFakeKlient(ctx)
	if err != nil {
		return err
	}

	zbiKlient, err := zklient.NewFakeZBIClient(klient)
	if err != nil {
		return err
	}

	k.klient = klient
	k.zbiKlient = zbiKlient

	return nil
}

func (k *FakeKlientFactory) GetZBIClient() interfaces.ZBIClientIF {
	return k.zbiKlient
}

func (k *FakeKlientFactory) GetKubernesClient() interfaces.KlientIF {
	return k.klient
}

func (k *FakeKlientFactory) StartInformer(ctx context.Context) error {
	logger.Infof(ctx, "Starting informer")
	k.zbiKlient.RunInformer(ctx)

	return nil
}
