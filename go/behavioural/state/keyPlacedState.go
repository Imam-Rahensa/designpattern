package main

import "fmt"

type keyPlacedState struct {
	config *config
}

func (u *keyPlacedState) updateKey() error {
	fmt.Println("Proposal key update is placed.")
	u.config.setState(u.config.keyPlaced)
	return nil
}

func (u *keyPlacedState) leadApproved() error {
	return fmt.Errorf("No proposal need to review.")
}

func (u *keyPlacedState) publishKey() error {
	return fmt.Errorf("No proposal need to approved.")
}
