package entity

type InitializeBkmResponse struct {
	Status         string `json:"status"`
	Locale         string `json:"locale"`
	SystemTime     int64  `json:"systemTime"`
	ConversationID string `json:"conversationId"`
	HtmlContent    string `json:"htmlContent"`
	Token          string `json:"token"`
	Checksum       string `json:"checksum"`
}

type CreatePaymentResponse struct {
	Status                    string            `json:"status"`
	Locale                    string            `json:"locale"`
	SystemTime                int64             `json:"systemTime"`
	ConversationID            string            `json:"conversationId"`
	Price                     float64           `json:"price"`
	PaidPrice                 float64           `json:"paidPrice"`
	Installment               int               `json:"installment"`
	PaymentID                 string            `json:"paymentId"`
	FraudStatus               int               `json:"fraudStatus"`
	MerchantCommissionRate    float64           `json:"merchantCommissionRate"`
	MerchantCommissionRateAmt float64           `json:"merchantCommissionRateAmount"`
	IyziCommissionRateAmt     float64           `json:"iyziCommissionRateAmount"`
	IyziCommissionFee         float64           `json:"iyziCommissionFee"`
	CardType                  string            `json:"cardType"`
	CardAssociation           string            `json:"cardAssociation"`
	CardFamily                string            `json:"cardFamily"`
	BinNumber                 string            `json:"binNumber"`
	LastFourDigits            string            `json:"lastFourDigits"`
	BasketID                  string            `json:"basketId"`
	Currency                  string            `json:"currency"`
	ItemTransactions          []ItemTransaction `json:"itemTransactions"`
	AuthCode                  string            `json:"authCode"`
	Phase                     string            `json:"phase"`
	HostReference             string            `json:"hostReference"`
	Checksum                  string            `json:"checksum"`
}

type ItemTransaction struct {
	ItemID                        string  `json:"itemId"`
	PaymentTransactionID          string  `json:"paymentTransactionId"`
	TransactionStatus             int     `json:"transactionStatus"`
	Price                         float64 `json:"price"`
	PaidPrice                     float64 `json:"paidPrice"`
	MerchantCommissionRate        float64 `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount  float64 `json:"merchantCommissionRateAmount"`
	IyziCommissionRateAmount      float64 `json:"iyziCommissionRateAmount"`
	IyziCommissionFee             float64 `json:"iyziCommissionFee"`
	BlockageRate                  float64 `json:"blockageRate"`
	BlockageRateAmountMerchant    float64 `json:"blockageRateAmountMerchant"`
	BlockageRateAmountSubMerchant float64 `json:"blockageRateAmountSubMerchant"`
	BlockageResolvedDate          string  `json:"blockageResolvedDate"`
	SubMerchantPrice              float64 `json:"subMerchantPrice"`
	SubMerchantPayoutRate         float64 `json:"subMerchantPayoutRate"`
	SubMerchantPayoutAmount       float64 `json:"subMerchantPayoutAmount"`
	MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
	ConvertedPayout               struct {
		PaidPrice                     float64 `json:"paidPrice"`
		IyziCommissionRateAmount      float64 `json:"iyziCommissionRateAmount"`
		IyziCommissionFee             float64 `json:"iyziCommissionFee"`
		BlockageRateAmountMerchant    float64 `json:"blockageRateAmountMerchant"`
		BlockageRateAmountSubMerchant float64 `json:"blockageRateAmountSubMerchant"`
		SubMerchantPayoutAmount       float64 `json:"subMerchantPayoutAmount"`
		MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
		IyziConversionRate            int     `json:"iyziConversionRate"`
		IyziConversionRateAmount      float64 `json:"iyziConversionRateAmount"`
		Currency                      string  `json:"currency"`
	} `json:"convertedPayout"`
}
