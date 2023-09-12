package rpcr

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/ariefsam/eventsam/idgenerator"
)

func (r RPCR) Call(endpoint string, input string) (output string, err error) {

	ctx := context.Background()
	inputRpc := Input{
		Endpoint:      endpoint,
		Input:         input,
		CorrelationID: idgenerator.Generate(),
	}

	// new json encoder allow html char
	w := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	err = encoder.Encode(inputRpc)
	if err != nil {
		log.Println(err)
		return
	}
	inputRpcString := w.String()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {

		respEndpoint := responseKey(inputRpc.CorrelationID, endpoint)
		timeout := 60 * time.Second
		res, err := r.rdb.BLPop(ctx, timeout, respEndpoint).Result()
		if err != nil {
			if err.Error() != "redis: nil" {
				log.Println(err)
			}
			return
		}
		output = res[1]
		wg.Done()
	}()

	_, err = r.rdb.RPush(ctx, endpoint, inputRpcString).Result()
	if err != nil {
		log.Println(err)
		return
	}
	wg.Wait()

	return
}
