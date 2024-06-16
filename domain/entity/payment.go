package entity

type Payment struct {
    Amount          float64
    Method          string
    Bank            string
    CardHolderName  string
    CardNumber      string
    ExpireMonth     string
    ExpireYear      string
    Cvc             string
    IdentityNumber  string
    BuyerEmail      string
    BuyerName       string
    BuyerSurname    string
    BuyerIp         string
    BuyerCity       string
    BuyerCountry    string
    BuyerAddress    string
    BuyerZipCode    string
}