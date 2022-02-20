package main

import (
	"ecosia/pkg/server"
	"log"
	"os"
)

func main() {

	if os.Getenv("IDLETIMEOUT") == "" || os.Getenv("READIMEOUT") == "" ||
		os.Getenv("WRITETIMEOUT") == "" {
		os.Setenv("IDLETIMEOUT", "1s")
		os.Setenv("READTIMEOUT", "1s")
		os.Setenv("WRITETIMEOUT", "1m")
	}

	srv := server.NewApplicationServer(9000, os.Getenv("IDLETIMEOUT"), os.Getenv("READTIMEOUT"),
		os.Getenv("WRITETIMEOUT"))
	err := srv.RunServer()
	if err != nil {
		log.Fatalln(err)
	}

}
