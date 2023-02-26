package base

import "encoding/json"

type Result struct {
	Data    interface{} `json:"data"`
	ErrCode int         `json:"err_code"`
	ErrMsg  string      `json:"error"`
}

type AliasResult Result

func (r Result) MarshalJSON() ([]byte, error) {
	if r.ErrMsg != "" {
		r.ErrCode = -1
	}
	if r.ErrMsg == "" && r.Data == nil {
		r.ErrMsg = "success"
	}
	return json.Marshal(AliasResult(r))
}
