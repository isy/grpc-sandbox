package ios

import (
	"context"

	"github.com/isy/grpc-sandbox/payment/app/usecase"
	pb "github.com/isy/grpc-sandbox/payment/pb/payment"
	"google.golang.org/protobuf/types/known/emptypb"
)

type (
	Ios interface {
		AppleVerifyReceipt(ctx context.Context, in *pb.AppleVerifyReceiptRequest) (*emptypb.Empty, error)
	}
	ios struct {
		uc usecase.Ios
	}
)

func NewIos(usecase usecase.Ios) Ios {
	return &ios{
		uc: usecase,
	}
}

func (i *ios) AppleVerifyReceipt(ctx context.Context, in *pb.AppleVerifyReceiptRequest) (*emptypb.Empty, error) {
	if err := i.uc.VerifyReceipt(ctx, in.ReceiptData); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
