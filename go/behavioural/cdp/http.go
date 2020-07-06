package main

import "github.com/nyelonong/designpattern-go/behavioural/cdp/toggle"

type http struct {
	command toggle.Toggle
}

func (h *http) press() {
	h.command.Execute()
}
