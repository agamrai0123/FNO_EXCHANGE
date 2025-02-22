package utils

import (
	"errors"

	"github.com/agamrai0123/FNO_EXCHANGE/ingest/models"
)

func ValidateOrderInputs(ord *models.Order) error {
	if ord.OrderFlow != 'B' && ord.OrderFlow != 'S' {
		return errors.New("invalid order flow")
	}
	if ord.SLMFlag != 'S' && ord.SLMFlag != 'L' && ord.SLMFlag != 'M' {
		return errors.New("invalid slm flag")
	}
	if ord.OrderType != 'I' && ord.OrderType != 'T' {
		return errors.New("order validity should be either 'Day' or 'IOC")
	}
	if ord.StopLossTrigger > 0 {
		ord.SLMFlag = 'S'
	}
	if ord.StopLossTrigger != 0 {
		if ord.OrderFlow == 'B' && ord.StopLossTrigger > ord.LimitRate {
			return errors.New("stop loss trigger price cannot be greater than limit rate")
		}
		if ord.OrderFlow == 'S' && ord.StopLossTrigger < ord.LimitRate {
			return errors.New("stop loss trigger price cannot be less than limit rate")
		}
	}
	if ord.SLMFlag == 'M' && ord.LimitRate != 0 {
		return errors.New("limit rate should be Zero for Market orders")
	}
	return nil
}
