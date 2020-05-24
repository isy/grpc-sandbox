package dao

import "github.com/isy/grpc-sandbox/payment/app/domain/repository"

type purchase struct{}

func NewPurchase() repository.Ios {
	return &purchase{}
}
