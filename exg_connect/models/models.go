package models

type GatewayRouterResponse struct {
	IPAddress        string
	Port             int32
	BoxId            int16
	SessionKey       [8]int8
	CryptographicKey []byte
	CryptographicIV  []byte
}

// type SocketInfo struct {
// 	Conn_type string
// 	Conn_host string
// 	Conn_port string
// 	Timeout   time.Duration
// }

type ExchangeData struct {
	Length         uint16
	SequenceNumber uint32
	Checksum       []byte
	MessageData    []byte
}
