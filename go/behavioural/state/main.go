package main

import "log"

func main() {
	flag := newConfig("feature-x", true)

	if err := flag.updateKey(); err != nil {
		log.Println(err)
	}

	if err := flag.publishKey(); err != nil {
		log.Println(err)
	}
}
