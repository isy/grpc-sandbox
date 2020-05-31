package repository

import (
	"context"

	"github.com/isy/grpc-sandbox/payment/app/domain/model"
)

type Ios interface {
	IosVerify(ctx context.Context, receipt string) (*model.Receipt, error)
}
