package main

import (
	"log"

	"github.com/spf13/cobra"
)

var mainCmd = &cobra.Command{
	Use:   "pingpong",
	Short: "pingpong docker networking test",
}

func main() {
	if err := mainCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
