package add

import "encoding/json"

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
