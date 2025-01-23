package main

import "os"

func openFile(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		return err
	}
	return err
}
