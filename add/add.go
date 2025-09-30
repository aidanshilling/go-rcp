package add

import (
	"encoding/json"
	"fmt"
	"log"
)

type AddMsg struct {
	ArgA int64 `json:"argA"`
	ArgB int64 `json:"argB"`
}

type AddResult struct {
	Result int64 `json:"result"`
}

func ReadAddResult(data []byte) AddResult {
	var res AddResult
	err := json.Unmarshal(data, &res)
	if err != nil {
		panic(err)
	}
	return res
}

func ReadAddMsg(data []byte) AddMsg {
	var msg AddMsg
	err := json.Unmarshal(data, &msg)
	if err != nil {
		panic(err)
	}

	return msg
}

func MarshalMsg(argA int64, argB int64) []byte {
	msg := AddMsg{ArgA: argA, ArgB: argB}
	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Failed to marshal message")
		log.Fatal(err)
	}

	return data
}

func MarshalAddResult(msg AddMsg) []byte {
	result := AddResult{Result: msg.ArgA + msg.ArgB}
	data, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	return data
}
