package ios

import (
	"context"

	"github.com/isy/grpc-sandbox/payment/app/usecase"
	pb_ios "github.com/isy/grpc-sandbox/payment/pb/payment/ios"
	"google.golang.org/protobuf/types/known/emptypb"
)

type (
	Ios interface {
		VerifyReceipt(ctx context.Context, in *pb_ios.VerifyReceiptRequest) (*emptypb.Empty, error)
	}
	ios struct {
		uc usecase.Ios
	}
)

func NewIos(usecase usecase.Ios) Ios {
	return &ios{}
}

func (i *ios) VerifyReceipt(ctx context.Context, in *pb_ios.VerifyReceiptRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
