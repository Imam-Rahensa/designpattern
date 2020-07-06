package main

import "fmt"

type keyApprovedState struct {
	config *config
}

func (u *keyApprovedState) updateKey() error {
	return fmt.Errorf("Can not update on review key.")
}

func (u *keyApprovedState) leadApproved() error {
	u.config.setState(u.config.keyApproved)
	fmt.Println("Proposal key update is approved.")
	return nil
}

func (u *keyApprovedState) publishKey() error {
	return fmt.Errorf("No proposal need to approved.")
}
