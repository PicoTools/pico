package errors

// GRPC errors
const (
	Send                   = "send"
	Receive                = "receive"
	GetFirstHelloMsg       = "get first message"
	NotHelloMsg            = "first message must be hello"
	SendListener           = "send listener"
	SendOperator           = "send operator"
	SendAgent              = "send agent"
	SendChatMessage        = "send chat message"
	SendCredential         = "send credential"
	SendCommandMessage     = "send command's message"
	SendCommandTask        = "send command's task"
	SendCommand            = "send command"
	UnauthenticatedRequest = "unauthenticated request"
	MissingRequestMetadata = "missing request metadata"
)

// logical errors
const (
	Internal                 = "internal error"
	VersionMismatched        = "client version mismatched"
	OperatorAlreadyConnected = "operator already connected"
	OperatorAlreadyExists    = "operator already exists"
	InvalidSessionCookie     = "invalid session cookie"
	GetSubscriptionData      = "unable get subscription data"
	DuringSubscription       = "something go wrong during subscription"
	AgentAlreadyExists       = "agent already exists"
	UnknownAgent             = "unknown agent"
	UnknownListener          = "unknown listener"
	UnknownOperator          = "unknown operator"
	UnknownTask              = "unknown task"
	CloseCommand             = "close command"
)

// parsing errors
const (
	MarshalCapability   = "marshal capability arguments"
	UnmarshalCapability = "unmarshal capability arguments"
	ParseIP             = "failed to parse IP"
	ParseExtIP          = "failed to parse ext IP"
	ParseIntIP          = "failed to parse int IP"
)

// database errors
const (
	BeginTx              = "begin tx"
	RollbackTx           = "rollback tx"
	CommitTx             = "commit tx"
	LoadOperator         = "load operator"
	QueryBlob            = "query blob"
	QueryListeners       = "query listeners"
	QueryListener        = "query listener"
	QueryOperators       = "query operators"
	QueryOperator        = "query operator"
	QueryAgents          = "query agents"
	QueryAgent           = "query agent"
	QueryChatMessages    = "query chat's messages"
	QueryChatMessage     = "query chat's message"
	QueryCredentials     = "query credentials"
	QueryCredential      = "query credential"
	QueryCommands        = "query commands"
	QueryCommandMessages = "query command's messages"
	QueryCommandTasks    = "query command's tasks"
	QueryCommandTask     = "query command's task"
	QueryPkiCa           = "query ca pki"
	QueryPkiOperator     = "query operator pki"
	QueryPkiListener     = "query listener pki"
	SaveChatMessage      = "save chat message"
	SaveCredential       = "save credential"
	SaveCommandMessage   = "save command's message"
	SaveCommandTask      = "save command's task"
	SaveCommand          = "save command"
	SaveOperator         = "save operator"
	SaveBlob             = "save blob"
	SaveAgent            = "save agent"
	SaveListener         = "save listener"
	UpdateListener       = "update listener"
	UpdateOperator       = "update operator"
	UpdateAgent          = "update agent"
	UpdateCredential     = "update credential"
	UpdateCommandTask    = "update command's task"
	UpdateLastAgent      = "update last checkout for agent"
	UpdateLastListener   = "update last checkout for listener"
)
