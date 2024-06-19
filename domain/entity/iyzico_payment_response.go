package entity

type IyzicoPaymentResponse struct {
	Status                       string `json:"status"`
	ErrorMessage                 string `json:"errorMessage"`
	ErrorCode                    string `json:"errorCode"`
	Locale                       string `json:"locale"`
	SystemTime                   int64  `json:"systemTime"`
	ConversationId               string `json:"conversationId"`
	PaymentId                    string `json:"paymentId"`
	FraudStatus                  int    `json:"fraudStatus"`
	Installment                  int    `json:"installment"`
	PaymentStatus                string `json:"paymentStatus"`
	Price                        string `json:"price"`
	PaidPrice                    string `json:"paidPrice"`
	MerchantCommissionRate       string `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount string `json:"merchantCommissionRateAmount"`
	IyziCommissionRateAmount     string `json:"iyziCommissionRateAmount"`
	IyziCommissionFee            string `json:"iyziCommissionFee"`
	CardType                     string `json:"cardType"`
	CardAssociation              string `json:"cardAssociation"`
	CardFamily                   string `json:"cardFamily"`
	CardToken                    string `json:"cardToken"`
	CardUserKey                  string `json:"cardUserKey"`
	BinNumber                    string `json:"binNumber"`
	BasketId                     string `json:"basketId"`
	Currency                     string `json:"currency"`
	ItemTransactions             []struct {
		ItemId                       string `json:"itemId"`
		PaymentTransactionId         string `json:"paymentTransactionId"`
		TransactionStatus            int    `json:"transactionStatus"`
		Price                        string `json:"price"`
		PaidPrice                    string `json:"paidPrice"`
		MerchantCommissionRate       string `json:"merchantCommissionRate"`
		MerchantCommissionRateAmount string `json:"merchantCommissionRateAmount"`
		IyziCommissionRateAmount     string `json:"iyziCommissionRateAmount"`
		IyziCommissionFee            string `json:"iyziCommissionFee"`
		BlockageRate                 string `json:"blockageRate"`
		BlockageRateAmountMerchant   string `json:"blockageRateAmountMerchant"`
		BlockageResolvedDate         string `json:"blockageResolvedDate"`
		SubMerchantPrice             string `json:"subMerchantPrice"`
		SubMerchantPayoutRate        string `json:"subMerchantPayoutRate"`
		SubMerchantPayoutAmount      string `json:"subMerchantPayoutAmount"`
		MerchantPayoutAmount         string `json:"merchantPayoutAmount"`
		ConvertedPayout              struct {
			PaidPrice                  string `json:"paidPrice"`
			IyziCommissionRateAmount   string `json:"iyziCommissionRateAmount"`
			IyziCommissionFee          string `json:"iyziCommissionFee"`
			BlockageRateAmountMerchant string `json:"blockageRateAmountMerchant"`
			MerchantPayoutAmount       string `json:"merchantPayoutAmount"`
			IyziConversionRate         string `json:"iyziConversionRate"`
			IyziConversionRateAmount   string `json:"iyziConversionRateAmount"`
		} `json:"convertedPayout"`
	} `json:"itemTransactions"`
}
type IyzicoInitializeResponse struct {
	Status         string `json:"status"`
	Locale         string `json:"locale"`
	SystemTime     int64  `json:"systemTime"`
	ConversationId string `json:"conversationId"`
	HtmlContent    string `json:"htmlContent"`
	Token          string `json:"token"`
	Checksum       string `json:"checksum"`
}
type RetrieveBkmResultResponse struct {
	Status         string `json:"status"`
	ErrorCode      string `json:"errorCode"`
	ErrorMessage   string `json:"errorMessage"`
	ErrorGroup     string `json:"errorGroup"`
	Locale         string `json:"locale"`
	SystemTime     int64  `json:"systemTime"`
	ConversationID string `json:"conversationId"`
}
type InitializeBkmResponse struct {
	Status         string `json:"status"`
	Locale         string `json:"locale"`
	SystemTime     int64  `json:"systemTime"`
	ConversationID string `json:"conversationId"`
	HtmlContent    string `json:"htmlContent"`
	Token          string `json:"token"`
	Checksum       string `json:"checksum"`
}
