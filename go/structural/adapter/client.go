package main

type client struct {
}

func (c *client) insertUsbInComputer(com computer) {
	com.insertUSB()
}
