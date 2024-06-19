package main

import "os"

func commandExit(_ *locationConfig, _ []string) error {
	os.Exit(0)
	return nil
}
