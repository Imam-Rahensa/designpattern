package main

func main() {
	client := &client{}
	mac := &mac{}
	client.insertUsbInComputer(mac)
	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowsMachine,
	}
	client.insertUsbInComputer(windowsMachineAdapter)
}
