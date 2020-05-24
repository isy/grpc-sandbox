package usecase

import (
	"context"

	"github.com/isy/grpc-sandbox/payment/app/domain/repository"
)

type (
	Ios interface {
		VerifyReceipt(ctx context.Context, receipt []byte, password string)
	}

	ios struct {
		ur repository.Ios
	}
)

func NewIos(ur repository.Ios) Ios {
	return &ios{
		ur: ur,
	}
}

func (i *ios) VerifyReceipt(ctx context.Context, receipt []byte, password string) {

}
