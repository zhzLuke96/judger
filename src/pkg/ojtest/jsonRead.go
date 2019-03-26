package ojtest

import (
	"encoding/json"

	"../utils"
)

func readTCaseFromJSONFile(filePth string) (caseBody testCaseConf, err error) {
	var fC []byte
	fC, err = utils.ReadFile(filePth)
	if err != nil {
		return caseBody, err
	}
	if err = json.Unmarshal(fC, &caseBody); err != nil {
		return caseBody, err
	}
	return caseBody, nil
}

func readTCaseFromJSONStr(jsonStr string) (caseBody testCaseConf, err error) {
	return caseBody, json.Unmarshal([]byte(jsonStr), &caseBody)
}
