package main

import (
	"fmt"

	"github.com/anxhukumar/gator-cli-tool/internal/config"
)

func main() {
	// read config file
	conf := config.Read()

	// set current user
	conf.SetUser("anxhukumar")

	// read config again and print the output
	updatedConfig := config.Read()
	fmt.Println(updatedConfig)
}
