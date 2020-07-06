package main

import "fmt"

type action struct {
	next ruleengine
}

func (a *action) execute(o *order) {
	if o.isActionDone {
		fmt.Println("Action already done")
		return
	}
	fmt.Println("Do action rule engine.")
	o.isActionDone = true
}

func (a *action) setNext(next ruleengine) {
	a.next = next
}
