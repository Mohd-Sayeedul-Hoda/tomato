package demon

import (
	"encoding/json"
)

type Method string

const (
	Status        Method = "STATUS"
	SessionCreate Method = "SESSION CREATE"
	NotDefine     Method = "NOT DEFINE"
)

type Request struct {
	Method Method          `json:"method"`
	Data   json.RawMessage `json:"data"`
}

type startTimerReq struct {
	SessionID  uint `json:"session_id"`
	Profile_id uint `json:"profile_id"`
}

type createSessionReq struct {
	Label    string `json:"label"`
	Note     string `json:"note"`
	Tracked  bool   `json:"tracked"`
	Estimate uint   `json:"estimate"`
}

type sessionListReq struct {
	Status  string `json:"status"`
	Tracked string `json:"tracked"`
}
