package apple

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/isy/grpc-sandbox/payment/app/infra/rest"
)

var (
	appleSharedSecret = os.Getenv("apple_shared_secret")
	ErrAppStoreServer = errors.New("AppStore server error")
)

func (a *AppStoreClient) ReqVerifyReceipt(ctx context.Context, receipt string) (*VerifyReceiptResponse, error) {
	body := &VerifyReceiptRequest{
		ReceiptData:            receipt,
		Password:               appleSharedSecret,
		ExcludeOldTransactions: true,
	}
	b, err := rest.EncodeBody(body)
	if err != nil {
		return nil, err
	}

	pReq, err := a.Prd.NewReq(ctx, "POST", "/verifyReceipt", b)
	if err != nil {
		return nil, fmt.Errorf("Error creating a client request to appstore in production: %w", err)
	}

	pRes, err := a.Prd.Do(pReq)
	if err != nil {
		return nil, fmt.Errorf("Appstore http request to production error: %w", err)
	}
	if pRes.StatusCode >= 500 {
		return nil, fmt.Errorf("Received http status code %d from the App Store: %w", pRes.StatusCode, ErrAppStoreServer)
	}

	resBody := &VerifyReceiptResponse{}
	if err := rest.DecodeBody(pRes, resBody); err != nil {
		return nil, err
	}

	// Receipt data for the sandbox
	if resBody.Status == 21007 {
		// In the case of receipt data for the sandbox, send a request to the sandbox environment.
		sReq, err := a.Sandbox.NewReq(ctx, "POST", "/verifyReceipt", b)
		if err != nil {
			return nil, fmt.Errorf("Error creating a client request to appstore in sandbox: %w", err)
		}

		sRes, err := a.Sandbox.Do(sReq)
		if err != nil {
			return nil, fmt.Errorf("Appstore http request to sandbox error: %w", err)
		}
		if sRes.StatusCode >= 500 {
			return nil, fmt.Errorf("Received http status code %d from the App Store Sandbox: %w", sRes.StatusCode, ErrAppStoreServer)
		}

		sResBody := &VerifyReceiptResponse{}
		if err := rest.DecodeBody(sRes, sResBody); err != nil {
			return nil, err
		}

		return sResBody, nil
	}

	return resBody, nil
}
