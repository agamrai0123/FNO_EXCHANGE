package utils

import (
	"github.com/agamrai0123/FNO_EXCHANGE/ingest/models"
	pb "github.com/agamrai0123/FNO_EXCHANGE/proto"
)

// ConvertProtoToModel converts a protobuf Order to your internal models.Order.
func ConvertProtoToModel(pbOrder *pb.Order) (models.Order, error) {
	var order models.Order

	order.SessionId = pbOrder.SessionId
	order.ExchangeCode = pbOrder.ExchangeCode
	order.EbaMatchAccount = pbOrder.EbaMatchAccount
	order.UserId = pbOrder.UserId
	order.Channel = pbOrder.Channel
	order.CseId = pbOrder.CseId
	order.PipeId = pbOrder.PipeId
	order.CtclId = pbOrder.CtclId

	// For fields stored as runes in your model, take the first character of the string (if available).
	if len(pbOrder.ProductType) > 0 {
		order.ProductType = rune(pbOrder.ProductType[0])
	}
	order.Underlying = pbOrder.Underlying
	order.ExpiryDate = pbOrder.ExpiryDate
	if len(pbOrder.ExcerciseType) > 0 {
		order.ExcerciseType = rune(pbOrder.ExcerciseType[0])
	}
	if len(pbOrder.OptionType) > 0 {
		order.OptionType = rune(pbOrder.OptionType[0])
	}

	// For numeric fields, cast as needed.
	order.StrikePrice = int32(pbOrder.StrikePrice)
	if len(pbOrder.IndexOrStock) > 0 {
		order.IndexOrStock = rune(pbOrder.IndexOrStock[0])
	}
	order.CALevel = pbOrder.CaLevel
	order.ActionId = pbOrder.ActionId
	order.BalanceAmount = pbOrder.BalanceAmount

	if len(pbOrder.CanModifyFlag) > 0 {
		order.CanModifyFlag = rune(pbOrder.CanModifyFlag[0])
	}
	if len(pbOrder.NkdBlockedFlag) > 0 {
		order.NKDBlockedFlag = rune(pbOrder.NkdBlockedFlag[0])
	}
	order.ModifyTradeDate = pbOrder.ModifyTradeDate
	order.ModifyTradeTime = pbOrder.ModifyTradeTime
	if len(pbOrder.SlmFlag) > 0 {
		order.SLMFlag = rune(pbOrder.SlmFlag[0])
	}

	order.DisclosedQuantity = int32(pbOrder.DisclosedQuantity)
	order.TotalOrderQuantity = int32(pbOrder.TotalOrderQuantity)
	order.LimitRate = int32(pbOrder.LimitRate)
	order.StopLossTrigger = int32(pbOrder.StopLossTrigger)
	order.OrderValidDate = pbOrder.OrderValidDate
	if len(pbOrder.OrderType) > 0 {
		order.OrderType = rune(pbOrder.OrderType[0])
	}
	order.AckTime = pbOrder.AckTime
	if len(pbOrder.SpecialFlag) > 0 {
		order.SpecialFlag = rune(pbOrder.SpecialFlag[0])
	}
	if len(pbOrder.OrderFlow) > 0 {
		order.OrderFlow = rune(pbOrder.OrderFlow[0])
	}
	if len(pbOrder.SpreadOrderIndicator) > 0 {
		order.SpreadOrderIndicator = rune(pbOrder.SpreadOrderIndicator[0])
	}
	order.Remarks = pbOrder.Remarks
	if len(pbOrder.UserFlag) > 0 {
		order.UserFlag = rune(pbOrder.UserFlag[0])
	}
	order.ExchangeRemarks = pbOrder.ExchangeRemarks
	order.IndexCode = pbOrder.IndexCode
	if len(pbOrder.SltpTrailFlag) > 0 {
		order.SLTPTrailFlag = rune(pbOrder.SltpTrailFlag[0])
	}
	order.VendorId = pbOrder.VendorId
	order.VendorNumber = pbOrder.VendorNumber
	if len(pbOrder.OneClickFlag) > 0 {
		order.OneClickFlag = rune(pbOrder.OneClickFlag[0])
	}
	order.OneClickPortfolioId = pbOrder.OneClickPortfolioId
	order.AlgoId = pbOrder.AlgoId
	order.AlgoOrderRemarks = pbOrder.AlgoOrderRemarks
	if len(pbOrder.SourceFlag) > 0 {
		order.SourceFlag = rune(pbOrder.SourceFlag[0])
	}
	if len(pbOrder.PopupFlag) > 0 {
		order.PopupFlag = rune(pbOrder.PopupFlag[0])
	}
	order.ExpiryDate2 = pbOrder.ExpiryDate2
	order.IpAddress = pbOrder.IpAddress
	order.CallSource = pbOrder.CallSource
	order.FreshOrderRef = pbOrder.FreshOrderRef
	order.Alias = pbOrder.Alias
	order.SystemMessage = pbOrder.SystemMessage
	if len(pbOrder.RequestType) > 0 {
		order.RequestType = rune(pbOrder.RequestType[0])
	}
	order.UserPassword = pbOrder.UserPassword
	if len(pbOrder.DeliveryEosFlag) > 0 {
		order.DeliveryEOSFlag = rune(pbOrder.DeliveryEosFlag[0])
	}
	order.OrderReference = pbOrder.OrderReference
	order.CoverOrderRef = pbOrder.CoverOrderRef

	return order, nil
}
