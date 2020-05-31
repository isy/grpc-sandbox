package usecase

import (
	"context"

	"github.com/isy/grpc-sandbox/payment/app/domain/repository"
)

type (
	Ios interface {
		VerifyReceipt(ctx context.Context, receipt string) error
	}

	ios struct {
		ri repository.Ios
	}
)

func NewIos(ri repository.Ios) Ios {
	return &ios{
		ri: ri,
	}
}

func (i *ios) VerifyReceipt(ctx context.Context, receipt string) error {
	_, err := i.ri.IosVerify(ctx, receipt)
	if err != nil {
		return err
	}

	return nil
}
