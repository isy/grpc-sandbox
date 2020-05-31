package apple

type (
	// Possible values: Sandbox, Production
	Environment          string
	OriginalPurchaseDate struct {
		OriginalPurchaseDate    string `json:"original_purchase_date"`
		OriginalPurchaseDateMs  string `json:"original_purchase_date_ms"`
		OriginalPurchaseDatePst string `json:"original_purchase_date_pst"`
	}

	// https://developer.apple.com/documentation/appstorereceipts/responsebody/receipt/in_app
	InApp struct {
		CancellationDate     string `json:"cancellation_date"`
		CancellationDateMs   string `json:"cancellation_date_ms"`
		CancellationDatePst  string `json:"cancellation_date_pst"`
		CancellationReason   string `json:"cancellation_reason"`
		ExpiresDate          string `json:"expires_date"`
		ExpiresDateMs        string `json:"expires_date_ms"`
		ExpiresDatePst       string `json:"expires_date_pst"`
		IsInIntroOfferPeriod string `json:"is_in_intro_offer_period"`
		IsTrialPeriod        string `json:"is_trial_period"`
		OriginalPurchaseDate
		OriginalTransactionID string `json:"original_transaction_id"`
		ProductID             string `json:"product_id"`
		PromotionalOfferID    string `json:"promotional_offer_id"`
		PurchaseDate          string `json:"purchase_date"`
		PurchaseDateMs        string `json:"purchase_date_ms"`
		PurchaseDatePst       string `json:"purchase_date_pst"`
		Quantity              string `json:"quantity"`
		TransactionID         string `json:"transaction_id"`
		WebOrderLineItemID    string `json:"web_order_line_item_id,omitempty"`
	}

	// https://developer.apple.com/documentation/appstorereceipts/responsebody/latest_receipt_info
	ReceiptInfo struct {
		InApp
		IsUpgraded                  string `json:"is_upgraded"`
		SubscriptionGroupIdentifier string `json:"subscription_group_identifier"`
	}

	// https://developer.apple.com/documentation/appstorereceipts/responsebody/pending_renewal_info
	PendingRenewalInfo struct {
		AutoRenewProductID     string `json:"auto_renew_product_id"`
		AutoRenewStatus        string `json:"auto_renew_status"`
		ExpirationIntent       string `json:"expiration_intent"`
		GracePeriodDate        string `json:"grace_period_expires_date"`
		GracePeriodDateMs      string `json:"grace_period_expires_date_ms"`
		GracePeriodDatePst     string `json:"grace_period_expires_date_pst"`
		IsInBillingRetryPeriod string `json:"is_in_billing_retry_period"`
		OriginalTransactionID  string `json:"original_transaction_id"`
		PriceConsentStatus     string `json:"price_consent_status"`
		ProductID              string `json:"product_id"`
	}

	// https://developer.apple.com/documentation/appstorereceipts/responsebody/receipt
	Receipt struct {
		AdamID                     int64   `json:"adam_id"`
		AppItemID                  int64   `json:"app_item_id"`
		ApplicationVersion         string  `json:"application_version"`
		BundleID                   string  `json:"bundle_id"`
		DownloadID                 int64   `json:"download_id"`
		ExpirationDate             string  `json:"expiration_date"`
		ExpirationDateMs           string  `json:"expiration_date_ms"`
		ExpirationDatePst          string  `json:"expiration_date_pst"`
		InApp                      []InApp `json:"in_app"`
		OriginalApplicationVersion string  `json:"original_application_version"`
		OriginalPurchaseDate
		PreorderDate           string `json:"preorder_date"`
		PreorderDateMs         string `json:"preorder_date_ms"`
		PreorderDatePst        string `json:"preorder_date_pst"`
		ReceiptCreationDate    string `json:"receipt_creation_date"`
		ReceiptCreationDateMs  string `json:"receipt_creation_date_ms"`
		ReceiptCreationDatePst string `json:"receipt_creation_date_pst"`
		// Possible values: Production, ProductionVPP, ProductionSandbox, ProductionVPPSandbox
		ReceiptType               string `json:"receipt_type"`
		RequestDate               string `json:"request_date"`
		RequestDateMs             string `json:"request_date_ms"`
		RequestDatePst            string `json:"request_date_pst"`
		VersionExternalIdentifier int64  `json:"version_external_identifier"`
	}

	// https://developer.apple.com/documentation/appstorereceipts/requestbody
	VerifyReceiptRequest struct {
		ReceiptData            string `json:"receipt-data"`
		Password               string `json:"password"`
		ExcludeOldTransactions bool   `json:"exclude-old-transactions,omitempty"`
	}

	// https://developer.apple.com/documentation/appstorereceipts/responsebody
	VerifyReceiptResponse struct {
		Environment        Environment          `json:"environment"`
		IsRetryable        bool                 `json:"is-retryable"`
		LatestReceipt      string               `json:"latest_receipt"`
		LatestReceiptInfo  []ReceiptInfo        `json:"latest_receipt_info"`
		PendingRenewalInfo []PendingRenewalInfo `json:"pending_renewal_info"`
		Receipt            Receipt              `json:"receipt"`
		// https://developer.apple.com/documentation/appstorereceipts/status
		Status int `json:"status"`
	}
)
