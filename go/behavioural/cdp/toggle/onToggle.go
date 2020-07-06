package toggle

import "github.com/nyelonong/designpattern-go/behavioural/cdp/flag"

type OnToggle struct {
	Flag flag.Flag
}

func (o *OnToggle) Execute() {
	o.Flag.On()
}
