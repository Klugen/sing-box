//go:build !go1.20

package shadowtls

import sTLS "github.com/sagernet/sing-box/transport/shadowtls/tls_go119"

type (
	sTLSConfig          = sTLS.Config
	sTLSConnectionState = sTLS.ConnectionState
	sTLSConn            = sTLS.Conn
)

var (
	sTLSCipherSuites = sTLS.CipherSuites
	sTLSClient       = sTLS.Client
)