package main

import "log"

func abortOnFail(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
