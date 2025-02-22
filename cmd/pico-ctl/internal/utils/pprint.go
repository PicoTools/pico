package utils

import (
	"fmt"
	"strings"
	"time"

	managementv1 "github.com/PicoTools/pico/pkg/proto/management/v1"
	"github.com/docker/go-units"
)

func PrettyOperator(o *managementv1.Operator) string {
	var s strings.Builder
	username := o.Username
	token := "[none]"
	last := "[never]"

	if o.Token != nil && o.Token.GetValue() != "" {
		token = o.Token.GetValue()
	}
	if !o.Last.AsTime().IsZero() {
		last = units.HumanDuration(time.Since(o.Last.AsTime()))
	}

	s.WriteString(fmt.Sprintf("%-10s %s\n", "Username:", username))
	s.WriteString(fmt.Sprintf("%-10s %s\n", "Token:", token))
	s.WriteString(fmt.Sprintf("%-10s %s", "Last:", last))
	return s.String()
}

func PrettyListener(l *managementv1.Listener) string {
	var s strings.Builder
	id := fmt.Sprintf("%d", l.Id)
	token := "[none]"
	name := "[none]"
	ip := "[none]"
	port := "[none]"
	last := "[none]"

	if l.Token != nil && l.Token.GetValue() != "" {
		token = l.Token.GetValue()
	}
	if l.Name != nil && l.Name.GetValue() != "" {
		name = l.Name.GetValue()
	}
	if l.Ip != nil && l.Ip.GetValue() != "" {
		ip = l.Ip.GetValue()
	}
	if l.Port != nil && l.Port.GetValue() != 0 {
		port = fmt.Sprintf("%d", l.Port.GetValue())
	}
	if !l.Last.AsTime().IsZero() {
		last = units.HumanDuration(time.Since(l.Last.AsTime()))
	}

	s.WriteString(fmt.Sprintf("%-10s %s\n", "ID:", id))
	s.WriteString(fmt.Sprintf("%-10s %s\n", "Token:", token))
	s.WriteString(fmt.Sprintf("%-10s %s\n", "Name:", name))
	s.WriteString(fmt.Sprintf("%-10s %s\n", "IP:", ip))
	s.WriteString(fmt.Sprintf("%-10s %s\n", "Port:", port))
	s.WriteString(fmt.Sprintf("%-10s %s", "Last:", last))
	return s.String()
}
