package dao

import (
	"context"

	"github.com/isy/grpc-sandbox/payment/app/domain/model"
	"github.com/isy/grpc-sandbox/payment/app/domain/repository"
	"github.com/isy/grpc-sandbox/payment/app/infra/rest/apple"
)

type purchase struct {
	a *apple.AppStoreClient
}

func NewIosPurchase(a *apple.AppStoreClient) repository.Ios {
	return &purchase{
		a: a,
	}
}

func (p *purchase) IosVerify(ctx context.Context, receipt string) (*model.Receipt, error) {
	v, err := p.a.ReqVerifyReceipt(ctx, receipt)
	if err != nil {
		return nil, err
	}

	return &model.Receipt{
		Status: v.Status,
	}, nil
}
