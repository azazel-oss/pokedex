package main

import "os"

func commandExit(_ *locationConfig) error {
	os.Exit(0)
	return nil
}
