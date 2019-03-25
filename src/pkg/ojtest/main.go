package ojtest

import (
	"errors"
	"fmt"
	"strings"

	"../confReader"
	"../utils"
)

type testDataSet [][]string

type testCaseConf struct {
	Timeout int         `json:"timeout"`
	Memory  int         `json:"mem"`
	Data    testDataSet `json:"data"`
}

// Call the program according to the use case and target language,
// and return the percentage of passing the test case
func RunTests(fileName string, casePth string, langType string) (per100 int, err error) {
	var (
		D       testCaseConf
		cmdText string
	)

	cmdText, err = confReader.GlobalConf.GetRunCmdWithActualFileName(langType, fileName)
	if err != nil {
		return 0, err
	}

	D, err = readTCaseFromJSON(casePth)

	if err != nil {
		return 0, err
	}

	return runWithTestDataWithTimeout(D.Timeout, D.Data, cmdText)
}

func runWithTestDataWithTimeout(timeout int, D testDataSet, CallText string) (per100 int, err error) {
	var passCount = 0

	for _, v := range D {
		actual := v[len(v)-1]
		INPUT := strings.Join(v[:len(v)-1], " ")

		// OUTPUT, stderr, err := utils.ShellCmdTimeoutWithStdin(timeout, CallText, INPUT)
		pres, err := utils.SetTimeoutExecCmdAndInput(CallText, INPUT, timeout)

		if err != nil {
			// return 0, err
			if err.Error() == "Timeout" {
				return 0, err
			}
			continue
		}
		if pres.Stderr != "" {
			return 0, fmt.Errorf("Runtime Error: %s", pres.Stderr)
		}

		// `\n88\n` <==> `88`
		OUTPUT := strings.TrimSpace(pres.Stdout)

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
