package ojtest

import (
	"encoding/json"
	"errors"
	"strings"

	"../confReader"
	"../utils"
)

type testDataSet [][]string

// Call the program according to the use case and target language, and return the percentage of passing the test case
func RunTests(fileName string, caseName string, langType string) (per100 int, err error) {
	var (
		D       testDataSet
		cmdText string
	)

	cmdText, err = confReader.GlobalConf.GetRunCmdWithActualFileName(langType, fileName)
	if err != nil {
		return 0, err
	}

	D, err = readTCaseFromJSON(caseName)
	if err != nil {
		return 0, err
	}

	return runWithTestData(D, cmdText)
}

func runWithTestData(D testDataSet, CallText string) (per100 int, err error) {
	var passCount = 0

	for _, v := range D {
		actual := v[len(v)-1]
		INPUT := strings.Join(v[:len(v)-1], " ")

		OUTPUT, err := utils.GetExecCmdOutput(CallText, INPUT)
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
		per100 = passCount * 100 / len(D)
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
