package main

import "fmt"

type keyPublishedState struct {
	config *config
}

func (u *keyPublishedState) updateKey() error {
	return fmt.Errorf("Can not update on review key.")
}

func (u *keyPublishedState) leadApproved() error {
	return fmt.Errorf("Key is already approved.")
}

func (u *keyPublishedState) publishKey() error {
	u.config.setState(u.config.keyPublished)
	fmt.Println("Key change published.")
	return nil
}
