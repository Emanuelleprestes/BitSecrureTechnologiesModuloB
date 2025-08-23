package main

import (
	"fmt"

	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/src/server"
)

func init() {
}

func main() {
	server := server.Newserver(":8080", nil)
	if err := server.Run(); err != nil {
		fmt.Println(err)
	}
}
