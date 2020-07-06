package flag

import "fmt"

type Key struct {
	isActive bool
}

func (t *Key) On() {
	t.isActive = true
	fmt.Println("Turning key flag on")
}

func (t *Key) Off() {
	t.isActive = false
	fmt.Println("Turning key flag off")
}
