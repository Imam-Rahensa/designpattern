package main

import "fmt"

type evaluator struct {
	next ruleengine
}

func (e *evaluator) execute(o *order) {
	if o.isEvaluateDone {
		fmt.Println("Evaluate already done")
		e.next.execute(o)
		return
	}
	fmt.Println("Evaluate order.")
	o.isEvaluateDone = true
	e.next.execute(o)
}

func (e *evaluator) setNext(next ruleengine) {
	e.next = next
}
