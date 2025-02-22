package models

type Order struct {
	SessionId            uint32  `json:"session_id"`
	ExchangeCode         string  `json:"exchange_code"`
	EbaMatchAccount      string  `json:"ebamatch_account"`
	UserId               string  `json:"user_id"`
	Channel              string  `json:"channel"`
	CseId                int32   `json:"cse_id"` //added in ROUT
	PipeId               string  `json:"pipe_id"`
	CtclId               string  `json:"ctcl_id"`
	ProductType          rune    `json:"product_type"`
	Underlying           string  `json:"underlying"`
	ExpiryDate           string  `json:"expiry_date"`
	ExcerciseType        rune    `json:"excercise_type"`
	OptionType           rune    `json:"option_type"`
	StrikePrice          int32   `json:"strike_price"`
	IndexOrStock         rune    `json:"index_or_stock"`
	CALevel              int32   `json:"ca_level"`
	ActionId             int32   `json:"action_id"`       // in spn_plc_ord error
	BalanceAmount        float32 `json:"balance_amount"`  // in spn_plc_ord error
	CanModifyFlag        rune    `json:"can_modify_flag"` // in spn_plc_ord error
	NKDBlockedFlag       rune    `json:"NKD_blocked_flag"`
	ModifyTradeDate      string  `json:"modify_trade_date"`
	ModifyTradeTime      string  `json:"modify_trade_time"`
	SLMFlag              rune    `json:"slm_flag"`
	DisclosedQuantity    int32   `json:"disclosed_quantity"`
	TotalOrderQuantity   int32   `json:"total_order_quantity"`
	LimitRate            int32   `json:"limit_rate"`
	StopLossTrigger      int32   `json:"stop_loss_trigger"`
	OrderValidDate       string  `json:"order_valid_date"`
	OrderType            rune    `json:"order_type"`
	AckTime              string  `json:"ack_time"`
	SpecialFlag          rune    `json:"special_flag"`
	OrderFlow            rune    `json:"order_flow"`
	SpreadOrderIndicator rune    `json:"spread_order_indicator"`
	Remarks              string  `json:"remarks"`
	UserFlag             rune    `json:"user_flag"`
	ExchangeRemarks      string  `json:"exchange_remarks"`
	IndexCode            string  `json:"index_code"` //pipe id
	SLTPTrailFlag        rune    `json:"sltp_trail_flag"`
	VendorId             string  `json:"vendor_id"`
	VendorNumber         string  `json:"venvendor_number"`
	OneClickFlag         rune    `json:"one_click_flag"`
	OneClickPortfolioId  string  `json:"one_click_portfolio_id"`
	AlgoId               string  `json:"algo_id"`
	AlgoOrderRemarks     string  `json:"algo_order_remarks"`
	SourceFlag           rune    `json:"source_flag"`
	PopupFlag            rune    `json:"popup_up"`
	ExpiryDate2          string  `json:"expiry_date_2"`
	IpAddress            string  `json:"ip_address"`
	CallSource           string  `json:"call_source"`
	FreshOrderRef        string  `json:"fresh_order_ref"` // profit_order
	Alias                string  `json:"alias"`
	SystemMessage        string  `json:"system_message"`
	RequestType          rune    `json:"request_type"`
	UserPassword         string  `json:"user_password"`
	DeliveryEOSFlag      rune    `json:"delivery_eos_flag"`
	OrderReference       string  `json:"order_reference"`
	CoverOrderRef        string  `json:"cover_order_ref"`
}
