package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "server",
	Short: "ping pong server",
	Run:   serverCommand,
}

func init() {
	mainCmd.AddCommand(serveCmd)
}

func serverCommand(cmd *cobra.Command, args []string) {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Println(time.Now(), ": received ping")
		fmt.Fprintf(w, "pong")
	})

	fmt.Println("Starting server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
