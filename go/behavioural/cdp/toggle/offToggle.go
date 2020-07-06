package toggle

import "github.com/nyelonong/designpattern-go/behavioural/cdp/flag"

type OffToggle struct {
	Flag flag.Flag
}

func (o *OffToggle) Execute() {
	o.Flag.Off()
}
