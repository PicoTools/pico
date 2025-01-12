package constants

import (
	"crypto/tls"
	"time"
)

const (
	// GrpcKeepalivePeriod is interval of QUIC keepalive messages
	GrpcKeepalivePeriod = 5 * time.Second
	// GrpcKeepaliveCount is count of QUIC keepalive mesages before connection marked as dead
	GrpcKeepaliveCount = 3
	// GrpcMaxConcurrentStreams is maximum number of concurrent streams. On limit reach new RPC calls
	// will be queue until capacity is available
	GrpcMaxConcurrentStreams = 500
	// GrpcTlsMinVersion is minimum TLS version for GRPC servers
	GrpcTlsMinVersion = tls.VersionTLS12
	// GrpcOperatorHealthCheckTimeout timeout to send health checks for operator
	GrpcOperatorHealthCheckTimeout = 10 * time.Second
)

var (
	// GrpcTlsCiphers is list of supported TLS ciphers
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
	// MaxObjectChunks is maximum number of objects in chunk.
	// Used for partiion data from DB on sub-arrays
	MaxObjectChunks int = 10000
)
