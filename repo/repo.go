package repo

import (
	"github.com/zbitech/common/interfaces"
	"github.com/zbitech/fake/repo/auth"
	"github.com/zbitech/fake/repo/service/factory"
)

func NewFakeServiceFactory() interfaces.ServiceFactoryIF {
	return factory.NewFakeServiceFactory()
}

func NewFakeAccessAuthorizerFactory() interfaces.AccessAuthorizerFactoryIF {
	return auth.NewFakeAuthorizerFactory()
}
