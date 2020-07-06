package main

type state interface {
	updateKey() error
	leadApproved() error
	publishKey() error
}
