package main

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/ariefsam/rpcr"
)

func main() {
	log.Default().SetFlags(log.LstdFlags | log.Llongfile)
	rpc := rpcr.NewRPCR()

	input := map[string]interface{}{
		"a": 1.1,
		"b": 2.3,
	}

	inputByte, _ := json.Marshal(input)
	inputString := string(inputByte)
	i := 0
	wg := sync.WaitGroup{}
	for {
		wg.Add(1)
		go func() {
			start := time.Now()
			output, err := rpc.Call("sum", inputString)
			end := time.Now()
			log.Println("Duration:", end.Sub(start))
			if err != nil {
				log.Println(err)
				return
			}
			log.Println(output)
			wg.Done()
		}()
		i++
		if i > 100 {
			break
		}
	}
	wg.Wait()
}
