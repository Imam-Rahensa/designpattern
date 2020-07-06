package main

type config struct {
	keyPlaced    state
	keyApproved  state
	keyPublished state

	currentState state

	key   string
	value bool
}

func newConfig(key string, value bool) *config {
	c := &config{
		key:   key,
		value: value,
	}

	keyPlacedState := &keyPlacedState{
		config: c,
	}

	keyApprovedState := &keyApprovedState{
		config: c,
	}

	keyPublishedState := &keyPublishedState{
		config: c,
	}

	c.setState(keyPlacedState)
	c.keyPlaced = keyPlacedState
	c.keyApproved = keyApprovedState
	c.keyPublished = keyPublishedState

	return c
}

func (c *config) updateKey() error {
	return c.currentState.updateKey()
}

func (c *config) leadApproved() error {
	return c.currentState.leadApproved()
}

func (c *config) publishKey() error {
	return c.currentState.publishKey()
}

func (c *config) setState(s state) {
	c.currentState = s
}
