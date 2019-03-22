package ojtest

import (
	"encoding/json"
	"errors"
	"strings"

	"../utils"
)

type testDataSet [][]string

func RunTests(fileName string, caseName string) (per100 int, err error) {
	var D testDataSet
	passCount := 0

	D, err = readTCaseFromJSON(caseName)
	if err != nil {
		return 0, err
	}
	for _, v := range D {
		actual := v[len(v)-1]
		INPUT := strings.Join(v[:len(v)-1], " ")

		OUTPUT, err := utils.GetExecCmdOutput(fileName, INPUT)
		if err != nil {
			return 0, err
		}

		// `\n88\n` <==> `88`
		OUTPUT = strings.TrimSpace(OUTPUT)

		if OUTPUT == actual {
			passCount++
		}
	}

	if passCount != len(D) {
		per100 = int((passCount / len(D)) * 100)
		return per100, errors.New("Testing case failed")
	}
	return 100, nil
}

func readTCaseFromJSON(filePth string) (D testDataSet, err error) {
	var fC []byte
	fC, err = utils.ReadFile(filePth)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(fC, &D); err != nil {
		return nil, err
	}
	return D, nil
}
