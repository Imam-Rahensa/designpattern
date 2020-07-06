package main

import (
	"github.com/nyelonong/designpattern-go/behavioural/cdp/flag"
	"github.com/nyelonong/designpattern-go/behavioural/cdp/toggle"
)

func main() {
	key := &flag.Key{}
	onToggle := &toggle.OnToggle{
		Flag: key,
	}
	offToggle := &toggle.OffToggle{
		Flag: key,
	}

	onFlag := &http{
		command: onToggle,
	}
	onFlag.press()

	offFlag := &nsq{
		command: offToggle,
	}
	offFlag.press()
}
