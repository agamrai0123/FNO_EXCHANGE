package models

// CHAR int8
// SHORT int16
// LONG int32
// UNSIGNED LONG uint32
// LONG LONG int64
// DOUBLE float64

// 40 bytes
type MESSAGE_HEADER struct {
	TransactionCode int16
	LogTime         int32
	AlphaChar       [2]int8
	TraderId        int32
	ErrorCode       int16
	Timestamp       int64
	TimeStamp1      [8]int8
	TimeStamp2      [8]int8
	MessageLength   int16
}

// 40 bytes
type INNER_MESSAGE_HEADER struct {
	TraderId        int32
	LogTime         int32
	AlphaChar       [2]int8
	TransactionCode int16
	ErrorCode       int16
	Timestamp       int64
	TimeStamp1      [8]int8
	TimeStamp2      [8]int8
	MessageLength   int16
}

// 40 bytes
type BCAST_HEADER struct {
	Reserved1       [2]int8
	Reserved2       [2]int8
	LogTime         int32
	AlphaChar       [2]int8
	TransactionCode int16
	ErrorCode       int16
	BCSeqNo         int32
	Reserved3       int8
	Reserved4       [3]int8
	TimeStamp2      [8]int8
	Filler2         [8]int8
	MessageLength   int16
}

// 182 bytes
type MS_ERROR_RESPONSE struct {
	MessageHeader MESSAGE_HEADER
	Key           [14]int8
	ErrorMessage  [128]int8
}

// MS_GR_REQUEST

// GR_REQUEST (2400)
// 48 bytes
type GR_REQUEST struct {
	MessageHeader MESSAGE_HEADER
	BoxId         int16
	BrokerId      [5]int8
	Filler        int8
}

// GR_RESPONSE(2401)
// 124 bytes
type GR_RESPONSE struct {
	MessageHeader    MESSAGE_HEADER
	BoxId            int16
	BrokerId         [5]int8
	Filler           int8
	IPAddress        [16]int8
	Port             int32
	SessionKey       [8]int8
	CryptographicKey [32]int8
	CryptographicIV  [16]int8
}

// MS_SECURE_BOX_REGISTRATION_REQUEST_IN

// SECURE_BOX_REGISTRATION_REQUEST_IN(23008)
// 42 bytes
type SECURE_BOX_REGISTRATION_REQUEST struct {
	MessageHeader MESSAGE_HEADER
	BoxId         int16
}

// SECURE_BOX_REGISTRATION_RESPONSE_OUT(23009)
// 40 bytes
type SECURE_BOX_REGISTRATION_RESPONSE struct {
	MessageHeader MESSAGE_HEADER
}

// MS_BOX_SIGN_ON_REQUEST_IN

// BOX_SIGN_ON_REQUEST_IN(23000)
// 60 bytes
type BOX_SIGN_ON_REQUEST_IN struct {
	MessageHeader MESSAGE_HEADER
	BoxId         int16
	BrokerId      [5]int8
	Reserved1     [5]int8
	SessionKey    [8]int8
}

// BOX_SIGN_ON_REQUEST_OUT(23001)
// 52 bytes
type BOX_SIGN_ON_REQUEST_OUT struct {
	MessageHeader MESSAGE_HEADER
	BoxId         int16
	Reserved1     [10]int8
}

// MS_BOX_SIGN_OFF

// BOX_SIGN_OFF (20322)
// 42 bytes
type BOX_SIGN_OFF struct {
	MessageHeader MESSAGE_HEADER
	BoxId         int16
}

// MS_SIGNON

// 2 bytes
type ST_BROKER_ELIGIBILITY_PER_MKT struct {
	MarketEligibilty uint8
	Reserved1        uint8
}

// SIGN_ON_REQUEST_IN (2300)
// 278 bytes
type SIGN_ON_REQUEST_IN struct {
	MessageHeader           MESSAGE_HEADER
	UserID                  int32
	Reserved1               [8]int8
	Passsword               [8]int8
	Reserved2               [8]int8
	NewPassword             [8]int8
	TraderName              [26]int8
	LastPasswordChangeDate  int32
	BrokerID                [5]int8
	Reserved3               int8
	BranchID                int16
	VersionNumber           int32
	Batch2StartTime         int32
	HostSwitchContext       int8
	Colour                  [50]int8
	Reserved4               int8
	UserType                int16
	SequenceNumber          float64
	WsClassName             [14]int8
	BrokerStatus            int8
	ShowIndex               int8
	BrokerEligibilityperMkt ST_BROKER_ELIGIBILITY_PER_MKT
	MemberType              int16
	ClearingStatus          int8
	Reserved5               [16]int8
	Reserved6               [16]int8
	Reserved7               [16]int8
}

// SIGN_ON_REQUEST_OUT (2301)
// 278 bytes
type SIGN_ON_REQUEST_OUT struct {
	MessageHeader           MESSAGE_HEADER
	UserID                  int32
	Reserved1               [8]int8
	Passsword               [8]int8
	Reserved2               [8]int8
	NewPassword             [8]int8
	TraderName              [26]int8
	LastPasswordChangeDate  int32
	BrokerID                [5]int8
	Reserved3               int8
	BranchID                int16
	VersionNumber           int32
	EndTime                 int32
	Reserved4               int8
	Colour                  [50]int8
	Reserved5               int8
	UserType                int16
	SequenceNumber          float64
	Reserved6               [14]int8
	BrokerStatus            int8
	ShowIndex               int8
	BrokerEligibilityPerMkt ST_BROKER_ELIGIBILITY_PER_MKT
	MemberType              int16
	ClearingStatus          int8
	BrokerName              [25]int8
	Reserved7               [16]int8
	Reserved8               [16]int8
	Reserved9               [16]int8
}

// 8 bytes
type ST_MARKET_STATUS struct {
	Normal  int16
	Oddlot  int16
	Spot    int16
	Auction int16
}

// 8 bytes
type ST_EX_MARKET_STATUS struct {
	Normal  int16
	Oddlot  int16
	Spot    int16
	Auction int16
}

// 8 bytes
type ST_PL_MARKET_STATUS struct {
	Normal  int16
	Oddlot  int16
	Spot    int16
	Auction int16
}

// 2 bytes
type ST_STOCK_ELIGIBLE_INDICATORS struct {
	StckElgblInd uint8
	Reserved1    uint8
}

// MS_SYSTEM_INFO_REQ

// SYSTEM_INFORMATION_IN (1600)
// 44 bytes
type SYSTEM_INFORMATION_IN struct {
	MessageHeader           MESSAGE_HEADER
	LastUpdatePortfolioTIme int32
}

// SYSTEM_INFORMATION_OUT(1601)
// 106 bytes
type SYSTEM_INFORMATION_OUT struct {
	MessageHeader                   MESSAGE_HEADER
	MarketStatus                    ST_MARKET_STATUS
	ExMarketStatus                  ST_EX_MARKET_STATUS
	PlMarketStatus                  ST_PL_MARKET_STATUS
	UpdatePortfolio                 int8
	MarketIndex                     int32
	DfltSttlmntPrdNormal            int16
	DfltSttlmntPrdSpot              int16
	DfltSttlmntPrdAuction           int16
	CompetitorPeriod                int16
	SolicitorPeriod                 int16
	WarningPercent                  int16
	VolumeFreezePercent             int16
	SnapQuoteTime                   int16
	Reserved1                       [2]int8
	BoardLotQuantity                int32
	TickSize                        int32
	MaximumGtcDays                  int16
	StockEligibleInd                ST_STOCK_ELIGIBLE_INDICATORS
	DisclosedQuantityPercentAllowed int16
	RiskFreeInterestRate            int32
}

// MS_UPDATE_LOCAL_DATABASE

// UPDATE_LOCALDB_IN(7300)
// 82 bytes
type UPDATE_LOCALDB_IN struct {
	MessageHeader             MESSAGE_HEADER
	LastUpdateSecurityTime    int32
	LastUpdateParticipantTime int32
	LastUpdateInstrumentTime  int32
	LastUpdateIndexTime       int32
	RequestForOpenOrders      int8
	Reserved1                 int8
	MarketStatus              ST_MARKET_STATUS
	ExMarketStatus            ST_EX_MARKET_STATUS
	PlMarketStatus            ST_PL_MARKET_STATUS
}

// UPDATE_LOCALDB_HEADER(7307)
// 42 bytes
type UPDATE_LOCALDB_HEADER struct {
	MessageHeader MESSAGE_HEADER
	Reserved1     [2]int8
}

// UPDATE_LOCALDB_TRAILER (7308)
// 42 bytes
type UPDATE_LOCALDB_TRAILER struct {
	MessageHeader MESSAGE_HEADER
	Reserved1     [2]int8
}

// MS_MESSAGE_DOWNLOAD

// DOWNLOAD_REQUEST (7000)
// 48 bytes
type DOWNLOAD_REQUEST struct {
	MessageHeader  MESSAGE_HEADER
	SequenceNumber float64
}

// SIGN_OFF_REQUEST_OUT(2321)
// 40 bytes
type SIGN_OFF_REQUEST_OUT struct {
	MessageHeader MESSAGE_HEADER
}

// Heartbeat(23506)
// 40 bytes
type HEARTBEAT struct {
	MessageHeader MESSAGE_HEADER
}
