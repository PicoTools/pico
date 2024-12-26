package types

import (
	"database/sql/driver"
	"fmt"
	"net"
)

// Inet stores underly net.IP type
type Inet struct {
	net.IP
}

// Scan implement Scanner interface
func (i *Inet) Scan(value any) (err error) {
	switch v := value.(type) {
	case nil:
	case []byte:
		if i.IP = net.ParseIP(string(v)); i.IP == nil {
			err = fmt.Errorf("invalid value of ip %q", v)
		}
	case string:
		if i.IP = net.ParseIP(v); i.IP == nil {
			err = fmt.Errorf("invalid value of ip %q", v)
		}
	default:
		err = fmt.Errorf("unexpected type %T", v)
	}
	return
}

// Value implements Valuer interface
func (i Inet) Value() (driver.Value, error) {
	return i.IP.String(), nil
}

// String returns string with IP
func (i Inet) String() string {
	if i.IP == nil {
		return ""
	}
	return i.IP.String()
}
