package models

type GatewayRouterResponse struct {
	IPAddress        string
	Port             int32
	BoxId            int16
	SessionKey       string
	CryptographicKey string
	CryptographicIV  string
}
