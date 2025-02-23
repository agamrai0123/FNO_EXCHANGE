package utils

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/spf13/viper"
)

func SocketConnect(config *viper.Viper, target string) (net.Conn, error) {
	conntype := config.GetString(target + ".CONN_TYPE")
	connip := config.GetString(target + ".CONN_HOST")
	connport := config.GetString(target + ".CONN_PORT")
	// timeout := config.GetInt(target + ".CONN_TIMEOUT")

	conn, err := net.DialTimeout(conntype, connip+":"+connport, 5*time.Second)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return nil, err
	}
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		log.Println("Failed to set TCP_NODELAY")
		return nil, err
	}
	if err := tcpConn.SetNoDelay(true); err != nil {
		log.Println("Error setting TCP_NODELAY:", err)
		return nil, err
	}
	return tcpConn, nil
}
func SocketWrite(buf *bytes.Buffer, conn net.Conn) error {
	_, err := conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println("Error sending message:", err)
		return err
	}
	return nil
}
func SocketRead(conn net.Conn, buf []byte) error {
	_, err := io.ReadFull(conn, buf)
	if err != nil {
		fmt.Println("failed to ReadFull:", err)
		return err
	}
	return nil
}

func LoadConfig() (*viper.Viper, error) {
	config := viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("config")
	config.AddConfigPath("./exg_connect/config")
	if err := config.ReadInConfig(); err != nil {
		log.Fatal("error on parsing configuration file", err)
		return nil, err
	}
	log.Println("config loaded successfully")
	return config, nil
}

// func configureTCPOptions(conn net.Conn) error {
// 	tcpConn, ok := conn.(*net.TCPConn)
// 	if !ok {
// 		return fmt.Errorf("connection is not a TCP connection")
// 	}

// 	// Disable Nagle's algorithm (TCP_NODELAY) to send small packets immediately.
// 	if err := tcpConn.SetNoDelay(true); err != nil {
// 		return fmt.Errorf("failed to set TCP_NODELAY: %w", err)
// 	}

// 	// Enable SO_KEEPALIVE to detect dead peers.
// 	if err := tcpConn.SetKeepAlive(true); err != nil {
// 		return fmt.Errorf("failed to enable keepalive: %w", err)
// 	}
// 	// Set keepalive period (example: 30 seconds).
// 	if err := tcpConn.SetKeepAlivePeriod(30 * time.Second); err != nil {
// 		return fmt.Errorf("failed to set keepalive period: %w", err)
// 	}

// 	// Linux-specific options.
// 	if runtime.GOOS == "linux" {
// 		rawConn, err := tcpConn.SyscallConn()
// 		if err != nil {
// 			return fmt.Errorf("failed to get raw connection: %w", err)
// 		}

// 		var sockOptErr error
// 		err = rawConn.Control(func(fd uintptr) {
// 			// TCP_QUICKACK: send ACKs immediately (Linux-specific).
// 			sockOptErr = syscall.SetsockoptInt(int(fd), syscall.IPPROTO_TCP, syscall.TCP_QUICKACK, 1)
// 			if sockOptErr != nil {
// 				return
// 			}
// 			// TCP_KEEPIDLE: idle time before sending keepalive probes (e.g., 30 seconds).
// 			sockOptErr = syscall.SetsockoptInt(int(fd), syscall.IPPROTO_TCP, syscall.TCP_KEEPIDLE, 30)
// 			if sockOptErr != nil {
// 				return
// 			}
// 			// TCP_KEEPINTVL: interval between keepalive probes (e.g., 10 seconds).
// 			sockOptErr = syscall.SetsockoptInt(int(fd), syscall.IPPROTO_TCP, syscall.TCP_KEEPINTVL, 10)
// 			if sockOptErr != nil {
// 				return
// 			}
// 			// TCP_KEEPCNT: number of keepalive probes before considering the connection dead (e.g., 3).
// 			sockOptErr = syscall.SetsockoptInt(int(fd), syscall.IPPROTO_TCP, syscall.TCP_KEEPCNT, 3)
// 		})
// 		if err != nil {
// 			return fmt.Errorf("error during raw connection control: %w", err)
// 		}
// 		if sockOptErr != nil {
// 			return fmt.Errorf("socket option error: %w", sockOptErr)
// 		}
// 	} else {
// 		log.Println("Linux-specific TCP options are not set because the OS is not Linux.")
// 	}

// 	return nil
// }
