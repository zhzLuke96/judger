package retMsg

import (
	"encoding/json"
	"fmt"
	"time"

	"../complier"
	"../ojtest"
	"../utils"
)

type RetBody struct {
	Status      string            `json:"status"`
	Message     string            `json:"message"`
	ErrorNumber int               `json:"errNo"`
	Details     map[string]string `json:"D"`

	IsWork  bool `json:"iswork"`
	Onerr   bool `json:"onerr"`
	Success bool `json:"success"`
	Score   int  `json:"score"`
}

func createRetBody(err error, passPer int, complied bool) *RetBody {
	ret := new(RetBody)
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
	ret.Details["version"] = utils.Version
	ret.Details["Date"] = time.Now().String()
	ret.Details["LOG"] = ""

	return ret
}

func (r *RetBody) DumpJSON() (JSONText string, err error) {
	if bytebuf, err := json.Marshal(r); err != nil {
		return "", err
	} else {
		return string(bytebuf), nil
	}
}

func (r *RetBody) PassMsg() string {
	if r.IsWork {
		return fmt.Sprintf("%d", r.Score)
	}
	return "0"
}

func Run(codeFilePth string, langType string, casePth string) (ret RetBody) {
	compliedProgramFilePth, err := complier.ComplieCodeFromFile(codeFilePth, langType)
	if err != nil {
		return *createRetBody(err, 0, false)
	}
	per100, err := ojtest.RunTests(compliedProgramFilePth, casePth, langType)
	return *createRetBody(err, per100, true)
}

func RunFromCtx(code string, langType string, testcase string) (ret RetBody) {
	compliedProgramFilePth, err := complier.ComplieCode(code, langType)
	if err != nil {
		return *createRetBody(err, 0, false)
	}
	per100, err := ojtest.RunTestsFromCaseString(compliedProgramFilePth, testcase, langType)
	return *createRetBody(err, per100, true)
}
