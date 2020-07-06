package main

import "github.com/nyelonong/designpattern-go/behavioural/cdp/toggle"

type nsq struct {
	command toggle.Toggle
}

func (h *nsq) press() {
	h.command.Execute()
}
