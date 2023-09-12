package main

import (
	"context"
	"encoding/json"
)

var sum = func(ctx context.Context, input string) (output string, err error) {

	var inputParse struct {
		A float64 `json:"a"`
		B float64 `json:"b"`
	}

	err = json.Unmarshal([]byte(input), &inputParse)
	if err != nil {
		return
	}

	outputParse := struct {
		Result float64 `json:"result"`
	}{
		Result: inputParse.A + inputParse.B,
	}

	outputByte, err := json.Marshal(outputParse)
	if err != nil {
		return
	}
	output = string(outputByte)

	return
}
