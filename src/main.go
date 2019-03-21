package main

import (
	"fmt"
	"strings"
)

var (
	complieCmd map[string]string = map[string]string{
		"py": "python3.exe -m py_compile "}
)

func complie(code string, env string) (fileName string, err error) {
	if fileName, err = saveStrAsFile(code); err != nil {
		return fileName, err
	}

	switch env {
	case "py":
		if _, err = getExecCmdOutput(complieCmd["py"]+"./"+fileName+".buf", ""); err != nil {
			return "", err
		}
		fileName = "python3.exe ./__pycache__/" + fileName + ".cpython-36.pyc"
		return fileName, nil
	default:
		fmt.Println("TODO!")
	}
	return
}

func runTests(fileName string, caseTable [][]string) (ok bool, err error) {
	for _, v := range caseTable {
		outstr := v[len(v)-1]
		instr := strings.Join(v[:len(v)-1], " ")
		tout, _ := callingOutputProgram(fileName, instr)
		tout = strings.TrimSpace(tout)
		if tout != outstr {
			return false, nil
		}
	}
	return true, nil
}

func callingOutputProgram(fileName string, input string) (output string, err error) {
	return getExecCmdOutput(fileName, input)
}

func main() {
	code, _ := ReadAll("adder.py")
	fn, err := complie(string(code), "py")
	if err != nil {
		fmt.Println(err)
	}

	testCase := [][]string{
		{"1+2", "3"},
		{"1-1", "0"},
		{"5-4", "1"},
		{"89+88", "177"}}

	if ok, err := runTests(fn, testCase); !ok {
		fmt.Println(err)
	}
	fmt.Println("Success!")

}
