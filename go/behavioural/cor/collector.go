package main

import "fmt"

type collector struct {
    next ruleengine
}

func (c *collector) execute(o *order) {
    if o.isCollectDone {
        fmt.Println("Collect already done")
        c.next.execute(o)
        return
    }
    fmt.Println("Collecting order data.")
	o.isCollectDone = true
    c.next.execute(o)
}

func (c *collector) setNext(next ruleengine) {
    c.next = next
}