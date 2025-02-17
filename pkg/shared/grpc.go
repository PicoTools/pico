package shared

const (
	// maximum length of single protobuf message (2GB)
	MaxProtobufMessageSize = 2 * 1024 * 1024 * 1024
	// Name of listener's authorization header
	GrpcAuthListenerHeader = "X-Access-Token"
	// Name of operator's authorization header
	GrpcAuthOperatorHeader = "Y-Access-Token"
	// Name of managements authorization header
	GrpcAuthManagementHeader = "Z-Access-Token"
)
