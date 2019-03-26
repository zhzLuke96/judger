package RPC

import (
	"../retMsg"
)

// Judger
type Judger struct{}

// Judger program Request
type JudgerRequest struct {
	Code         string
	LangType     string
	TestCaseJSON string
}

// Judger Ret
type JudgerResponse struct {
	JSONContent string
}

// Judger Work
func (this *Judger) DoJudger(req JudgerRequest, res *JudgerResponse) (err error) {
	result := retMsg.RunFromCtx(req.Code, req.LangType, req.TestCaseJSON)
	res.JSONContent, err = result.DumpJSON()
	if err != nil {
		return err
	}
	return nil
}
