package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type RetMsgBody struct {
	Status      string            `json:"status"`
	Message     string            `json:"message"`
	ErrorNumber int               `json:"errNo"`
	Details     map[string]string `json:"D"`

	IsWork  bool `json:"iswork"`
	Onerr   bool `json:"onerr"`
	Success bool `json:"success"`
	Score   int  `json:"score"`
}

func createMsgBody(err error, passPer int, complied bool) *RetMsgBody {
	ret := new(RetMsgBody)
	ret.IsWork, ret.Success, ret.Onerr = false, false, false
	ret.Score = passPer
	ret.Details = make(map[string]string)

	if err != nil {
		ret.Onerr = true
		// ret.Success = false
	}
	// else {
	// 	ret.Onerr = false
	// }
	if passPer != 0 {
		ret.IsWork = true
	}
	// else {
	// 	ret.Passed = false
	// }
	if passPer == 100 {
		ret.Success = true
	}
	// else {
	// 	ret.Success = false
	// }

	// Status
	switch {
	case err != nil && err.Error() == "Timeout":
		ret.Status = "Runtime Timeout"
	case !complied && ret.Onerr:
		ret.Status = "Complie ERROR"
	case complied && ret.Onerr:
		ret.Status = "Runtime ERROR"
	case ret.Success:
		ret.Status = "Success"
	}

	// Msg
	if err != nil {
		ret.Message = err.Error()
	} else {
		ret.Message = fmt.Sprintf("Passed %d%%", passPer)
	}

	// ErrorNumber
	// [TODO]
	ret.ErrorNumber = -1

	// Details
	ret.Details["version"] = _Version
	ret.Details["Date"] = time.Now().String()
	ret.Details["LOG"] = ""

	return ret
}

func (r *RetMsgBody) DumpJSON() (JSONText string, err error) {
	if bytebuf, err := json.Marshal(r); err != nil {
		return "", err
	} else {
		return string(bytebuf), nil
	}
}

func (r *RetMsgBody) PassMsg() string {
	if r.IsWork {
		return fmt.Sprintf("%d", r.Score)
	}
	return "0"
}
