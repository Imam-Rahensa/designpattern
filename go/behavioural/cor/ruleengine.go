package main

type ruleengine interface {
	execute(*order)
	setNext(ruleengine)
}
