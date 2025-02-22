package internal

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/agamrai0123/FNO_EXCHANGE/random_order_generator/models"
)

// Constants for generating random values
var (
	exchangeCodes = []string{"NSE", "BSE", "MCX", "NFO", "BFO"}
	channel       = []string{"CNT", "SYS", "WEB", "OFF"}
	underlyings   = []string{"NIFTY", "BANKNIFTY", "RELIANCE", "TCS", "INFY"}
	callSources   = []string{"WEB", "MOBILE", "API"}
	vendorIds     = []string{"V001", "V002", "V003"}
)

// Helper function to generate random string of given length
func randomString(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func randomNumber(length int) string {
	const charset = "0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// Helper function to generate random date within last 30 days
func randomDate() string {
	days := rand.Intn(30)
	date := time.Now().AddDate(0, 0, -days)
	return date.Format("2006-01-02")
}

// Helper function to generate random time
func randomTime() string {
	hour := rand.Intn(8) + 9 // 9 AM to 4 PM
	minute := rand.Intn(60)
	second := rand.Intn(60)
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}

// Helper function to generate random rune from given options
func randomRune(options string) rune {
	return rune(options[rand.Intn(len(options))])
}

// GenerateRandomOrder creates a new models.Order with random values
func GenerateRandomOrder() models.Order {

	order := models.Order{
		SessionId:            uint32(rand.Intn(99999) + 1),
		ExchangeCode:         exchangeCodes[rand.Intn(len(exchangeCodes))],
		EbaMatchAccount:      fmt.Sprintf("EBA%s", randomNumber(10)),
		UserId:               fmt.Sprintf("USER%s", randomString(6)),
		Channel:              channel[rand.Intn(len(channel))],
		CseId:                int32(rand.Intn(99999) + 1),
		PipeId:               fmt.Sprintf("PIPE%s", randomString(4)),
		CtclId:               fmt.Sprintf("CTCL%s", randomString(4)),
		ProductType:          randomRune("FO"),
		Underlying:           underlyings[rand.Intn(len(underlyings))],
		ExpiryDate:           randomDate(),
		ExcerciseType:        randomRune("EA"),
		OptionType:           randomRune("CPX"),
		StrikePrice:          int32(rand.Intn(1000)+100) * 100, // Random strike price between 100 and 100000
		IndexOrStock:         randomRune("IS"),
		CALevel:              int32(rand.Intn(99999) + 1),
		ActionId:             int32(rand.Intn(99999) + 1),
		BalanceAmount:        float32(rand.Float64() * 100000),
		CanModifyFlag:        randomRune("YN"),
		NKDBlockedFlag:       randomRune("YN"),
		ModifyTradeDate:      randomDate(),
		ModifyTradeTime:      randomTime(),
		SLMFlag:              randomRune("SLM"),
		DisclosedQuantity:    int32(rand.Intn(100) + 1),
		TotalOrderQuantity:   int32(rand.Intn(1000) + 1),
		LimitRate:            int32(rand.Intn(10000) + 100),
		StopLossTrigger:      int32(rand.Intn(10000) + 100),
		OrderValidDate:       randomDate(),
		OrderType:            randomRune("IT"),
		AckTime:              randomTime(),
		SpecialFlag:          randomRune("YN"),
		OrderFlow:            randomRune("BS"),
		SpreadOrderIndicator: randomRune("*"),
		Remarks:              fmt.Sprintf("Test Order %s", randomString(4)),
		UserFlag:             randomRune("YN"),
		ExchangeRemarks:      fmt.Sprintf("Exchange Remarks %s", randomString(4)),
		IndexCode:            fmt.Sprintf("IDX%s", randomString(3)),
		SLTPTrailFlag:        randomRune("YN"),
		VendorId:             vendorIds[rand.Intn(len(vendorIds))],
		VendorNumber:         fmt.Sprintf("VN%s", randomString(6)),
		OneClickFlag:         randomRune("YN"),
		OneClickPortfolioId:  fmt.Sprintf("PORT%s", randomString(4)),
		AlgoId:               fmt.Sprintf("ALGO%s", randomString(4)),
		AlgoOrderRemarks:     fmt.Sprintf("Algo Remarks %s", randomString(4)),
		SourceFlag:           randomRune("WMA"), // Web, Mobile, API
		PopupFlag:            randomRune("YN"),
		ExpiryDate2:          randomDate(),
		IpAddress:            fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256)),
		CallSource:           callSources[rand.Intn(len(callSources))],
		FreshOrderRef:        fmt.Sprintf("FOR%s", randomString(6)),
		Alias:                fmt.Sprintf("ALIAS%s", randomString(4)),
		SystemMessage:        fmt.Sprintf("SYS_MSG_%s", randomString(8)),
		RequestType:          randomRune("NMCT"), // New, Modify, Cancel, Trade
		UserPassword:         randomString(12),
		DeliveryEOSFlag:      randomRune("YN"),
		OrderReference:       fmt.Sprintf("REF%s", randomString(8)),
		CoverOrderRef:        fmt.Sprintf("COR%s", randomString(8)),
	}

	return order
}

// GenerateRandomOrders creates a slice of random orders
func GenerateRandomOrders(count int) []models.Order {
	orders := make([]models.Order, count)
	for i := 0; i < count; i++ {
		orders[i] = GenerateRandomOrder()
	}
	return orders
}
