package utils

import (
	"fmt"
	"time"

	managementv1 "github.com/PicoTools/pico/pkg/proto/management/v1"
	"github.com/docker/go-units"
	"github.com/jedib0t/go-pretty/v6/table"
)

// TableOperators prints table with operators
func TableOperators(operators []*managementv1.Operator) {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"Username", "Token", "Last"})
	for _, operator := range operators {
		t.AppendRow(tableOperatorRow(operator))
	}
	fmt.Println(t.Render())
}

// TableOperator prints table with operator
func TableOperator(operator *managementv1.Operator) {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"Username", "Token", "Last"})
	t.AppendRow(tableOperatorRow(operator))
	fmt.Println(t.Render())
}

// tableOperatorRow returns single row with operator's data for table
func tableOperatorRow(operator *managementv1.Operator) table.Row {
	username := operator.Username
	token := "[none]"
	last := "[never]"

	if operator.Token != nil && operator.Token.GetValue() != "" {
		token = operator.Token.GetValue()
	}
	if !operator.Last.AsTime().IsZero() {
		last = units.HumanDuration(time.Since(operator.Last.AsTime()))
	}

	return table.Row{username, token, last}
}

// TableListeners prints table with listener
func TableListeners(listeners []*managementv1.Listener) {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"ID", "Token", "Name", "IP", "Port", "Last"})
	for _, listener := range listeners {
		t.AppendRow(tableListenerRow(listener))
	}
	fmt.Println(t.Render())
}

// TableOperator prints table with listener
func TableListener(listener *managementv1.Listener) {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"ID", "Token", "Name", "IP", "Port", "Last"})
	t.AppendRow(tableListenerRow(listener))
	fmt.Println(t.Render())
}

// tableListenerRow returns single row with listener's data for table
func tableListenerRow(listener *managementv1.Listener) table.Row {
	id := fmt.Sprintf("%d", listener.Id)
	token := "[none]"
	name := "[none]"
	ip := "[none]"
	port := "[none]"
	last := "[none]"

	if listener.Token != nil && listener.Token.GetValue() != "" {
		token = listener.Token.GetValue()
	}
	if listener.Name != nil && listener.Name.GetValue() != "" {
		name = listener.Name.GetValue()
	}
	if listener.Ip != nil && listener.Ip.GetValue() != "" {
		ip = listener.Ip.GetValue()
	}
	if listener.Port != nil && listener.Port.GetValue() != 0 {
		port = fmt.Sprintf("%d", listener.Port.GetValue())
	}
	if !listener.Last.AsTime().IsZero() {
		last = units.HumanDuration(time.Since(listener.Last.AsTime()))
	}

	return table.Row{id, token, name, ip, port, last}
}
