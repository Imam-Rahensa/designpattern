package main

type windowsAdapter struct {
	windowMachine *windows
}

func (w *windowsAdapter) insertUSB() {
	w.windowMachine.insertInCirclePort()
}
