package shared

const (
	// default color used in process of object's creation
	DefaultObjectColor = uint32(0)
)

const (
	// maximum length of operator's username
	OperatorUsernameMaxLength = 256
	// minimum length of operator's username
	OperatorUsernameMinLength = 3
)

const (
	// maximum length of listener' name
	ListenerNameMaxLength = 256
	// maximum length of listener's note
	ListenerNoteMaxLength = 256
)

const (
	// maximum length of ant's operating system meta information
	AntOsMetaMaxLength = 1024
	// maximum length of ant's hostname
	AntHostnameMaxLength = 256
	// maximum length of ant's username
	AntUsernameMaxLength = 256
	// maximum length of ant's domain
	AntDomainMaxLength = 256
	// maximum length of ant's process name
	AntProcessNameMaxLength = 1024
	// maximum length of ant's note
	AntNoteMaxLength = 256
)

const (
	// minimum length of chat's message
	ChatMessageMinLength = 1
	// maximum length of chat's message
	ChatMessageMaxLength = 4096
)

const (
	// maximum length of credential's username
	CredentialUsernameMaxLength = 256
	// maximum length of credential's secret
	CredentialSecretMaxLength = 4096
	// maximum length of credential's realm
	CredentialRealmMaxLength = 256
	// maximum length of credential's host
	CredentialHostMaxLength = 256
	// maximum length of credential's note
	CredentialNoteMaxLength = 256
)

const (
	// minimum length of task group command line
	TaskGroupCmdMinLength = 1
	// maximum length of task group command line
	TaskGroupCmdMaxLength = 4096
	// maximum length of message in task group
	TaskGroupMessageMaxLength = 4096
)
