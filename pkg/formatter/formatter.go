package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"payment-system/domain/entity"
	"strings"
)

func ToJSON(obj interface{}) string {
	jsonObject, _ := json.Marshal(obj)
	log.Println(string(jsonObject))
	return string(jsonObject)
}
func FormatBasketItems(items []entity.BasketItem) string {
	var result []string
	for _, item := range items {
		itemStr := fmt.Sprintf("[id=%s,price=%s,name=%s,category1=%s,category2=%s,itemType=%s]",
			item.Id, item.Price, item.Name, item.Category1, item.Category2, item.ItemType)
		result = append(result, itemStr)
	}
	return strings.Join(result, ", ")
}
func FormatAddress(address entity.Address) string {
	return fmt.Sprintf("[address=%s,zipCode=%s,contactName=%s,city=%s,country=%s]",
		address.Address, address.ZipCode, address.ContactName, address.City, address.Country)
}
func FormatBuyer(buyer entity.Buyer) string {
	return fmt.Sprintf("[id=%s,name=%s,surname=%s,identityNumber=%s,email=%s,gsmNumber=%s,registrationDate=%s,lastLoginDate=%s,registrationAddress=%s,city=%s,country=%s,zipCode=%s,ip=%s]",
		buyer.Id, buyer.Name, buyer.Surname, buyer.IdentityNumber, buyer.Email, buyer.GsmNumber, buyer.RegistrationDate, buyer.LastLoginDate, buyer.RegistrationAddress, buyer.City, buyer.Country, buyer.ZipCode, buyer.Ip)
}

func FormatInitializeBkm(initializeBkm entity.InitializeBkmRequest) string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	buffer.WriteString(fmt.Sprintf("locale=%s,", initializeBkm.Locale))
	buffer.WriteString(fmt.Sprintf("conversationId=%s,", initializeBkm.ConversationID))
	buffer.WriteString(fmt.Sprintf("price=%s,", initializeBkm.Price))
	buffer.WriteString(fmt.Sprintf("basketId=%s,", initializeBkm.BasketID))
	buffer.WriteString(fmt.Sprintf("paymentGroup=%s,", initializeBkm.PaymentGroup))
	buffer.WriteString("buyer=")
	buffer.WriteString(FormatBuyer(initializeBkm.Buyer))
	buffer.WriteString(",")
	buffer.WriteString("shippingAddress=")
	buffer.WriteString(FormatAddress(initializeBkm.ShippingAddress))
	buffer.WriteString(",")
	buffer.WriteString("billingAddress=")
	buffer.WriteString(FormatAddress(initializeBkm.BillingAddress))
	buffer.WriteString(",")
	buffer.WriteString("basketItems=[")
	buffer.WriteString(FormatBasketItems(initializeBkm.BasketItems))
	buffer.WriteString("],")
	buffer.WriteString(fmt.Sprintf("callbackUrl=%s", initializeBkm.CallbackURL))
	buffer.WriteString("]")

	return buffer.String()
}
