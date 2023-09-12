package main

import (
	"log"
	"time"

	"github.com/ariefsam/rpcr"
)

func main() {
	log.Default().SetFlags(log.LstdFlags | log.Llongfile)
	server := rpcr.NewRPCR()
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		go func() {
			server.Serve("sum", sum)
		}()
	}
	server.Serve("sum", sum)
}
