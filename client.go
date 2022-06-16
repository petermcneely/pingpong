package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "ping pong client",
	Run:   clientCommand,
}

func init() {
	mainCmd.AddCommand(clientCmd)
}

func clientCommand(cmd *cobra.Command, args []string) {
	pingpongaddr := os.Getenv("PING_PONG_ADDR")
	t := time.NewTicker(time.Second)
	for now := range t.C {
		resp, err := http.Get(pingpongaddr + "/ping")
		if err != nil {
			log.Println(err.Error())
			break
		}

		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Println(err.Error())
		} else {
			log.Println(now, string(body))
		}
	}
}
