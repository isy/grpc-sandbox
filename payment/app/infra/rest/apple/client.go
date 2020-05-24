package apple

import (
	"net/url"

	"github.com/isy/grpc-sandbox/payment/app/infra/rest"
)

const (
	storePrd     = "https://buy.itunes.apple.com/verifyReceipt"
	storeSandbox = "https://sandbox.itunes.apple.com/verifyReceipt"
)

type AppStoreClient struct {
	Prd     *rest.Client
	Sandbox *rest.Client
}

func NewAppStoreClient() (*AppStoreClient, error) {
	pu, err := url.Parse(storePrd)
	if err != nil {
		return nil, err
	}
	prd := rest.NewHttpClient(pu)

	su, err := url.Parse(storeSandbox)
	if err != nil {
		return nil, err
	}
	sandbox := rest.NewHttpClient(su)

	return &AppStoreClient{
		Prd:     prd,
		Sandbox: sandbox,
	}, nil
}
