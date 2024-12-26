package constants

import (
	"crypto/tls"
	"time"
)

const (
	// send keepalive message every 30 seconds
	GrpcKeepaliveTime = time.Second * 30
	// keepalive timeout is 5 seconds
	GrpcKeepaliveTimeout = time.Second * 5
	// timeout after which connection will be gracefull closed
	GrpcMaxConnAgeGrace = time.Second * 10
	// minimum TLS version for GRPC servers
	GrpcTlsMinVersion = tls.VersionTLS12
	// timeout to send health checks for operator
	GrpcOperatorHealthCheckTimeout = time.Second * 10
)

var (
	// supported TLS ciphers
	GrpcTlsCiphers = []uint16{
		tls.TLS_AES_256_GCM_SHA384,
		tls.TLS_CHACHA20_POLY1305_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
	}
)

const (
	// maximum number of objects in chunk. Used for partiion data from DB on sub-arrays
	MaxObjectChunks int = 10000
)
