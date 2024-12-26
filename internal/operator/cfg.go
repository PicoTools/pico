package operator

import (
	"net"
)

type Config struct {
	IP   net.IP `json:"ip" validate:"required"`
	Port int    `json:"port" validate:"required,port"`
}
