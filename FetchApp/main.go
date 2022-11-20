package main

import (
	"fmt"
	"log"
	"os"
)

var (
	address string
)

func init() {
	address = fmt.Sprintf("0.0.0.0:%s", os.Getenv("port"))
}

func main() {
	log.Fatal(RunServer(address))
}
