package listener

import (
	"net"
)

// listener's server config
type Config struct {
	IP   net.IP `json:"ip" validate:"required"`
	Port int    `json:"port" validate:"required,port"`
}
