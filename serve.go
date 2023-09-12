package rpcr

import (
	"context"
	"encoding/json"
	"log"
	"time"
)

func (r RPCR) Serve(endpoint string, fnc func(ctx context.Context, input string) (output string, err error)) {
	for {
		r.serve(endpoint, fnc)
	}

}

func (r RPCR) serve(endpoint string, fnc func(ctx context.Context, input string) (output string, err error)) {
	timeout := 60 * time.Second
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	res, err := r.rdb.BLPop(ctx, timeout, endpoint).Result()

	if err != nil {
		log.Println(err)
		return
	}

	input := Input{}

	err = json.Unmarshal([]byte(res[1]), &input)
	if err != nil {
		log.Println(err)
		return
	}

	output, err := fnc(ctx, input.Input)
	if err != nil {
		log.Println(err)
		return
	}

	responseKey := responseKey(input.CorrelationID, endpoint)

	_, err = r.rdb.RPush(ctx, responseKey, output).Result()
	if err != nil {
		log.Println(err)
		return
	}

	r.rdb.Expire(ctx, responseKey, 5*time.Second).Result()

}
