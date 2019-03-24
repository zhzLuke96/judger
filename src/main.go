package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"./pkg/complier"
	"./pkg/confReader"
	"./pkg/ojtest"
)

const (
	confFilePth = "./conf.json"
	_Version    = "a0.1"
)

func clearCompliceFile() (err error) {
	os.RemoveAll("./__pycache__/")
	return nil
}

func run(codeFilePth string, langType string, casePth string) (ret RetMsgBody) {
	programFilePth, err := complier.ComplieCodeFromFile(codeFilePth, langType)
	if err != nil {
		return *createMsgBody(err, 0, false)
	}
	per100, err := ojtest.RunTests(programFilePth, casePth, langType)
	return *createMsgBody(err, per100, true)
}

func main() {
	var (
		langType   string
		casePth    string
		filePth    string
		outputMode string
	)
	flag.StringVar(&langType, "lang", "cpp", "Target language type of input code.")
	flag.StringVar(&langType, "l", "cpp", "Target language type of input code.")

	flag.StringVar(&casePth, "case", "0x1", "the file path of the test case.")
	flag.StringVar(&casePth, "c", "0x1", "the file path of the test case.")

	flag.StringVar(&outputMode, "out", "pure", "choice program output style [pure,json,onlyPass].")
	flag.StringVar(&outputMode, "o", "pure", "choice program output style [pure,json,onlyPass].")

	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("flag needs an argument for fileName.")
		flag.Usage()
		return
	}
	filePth = flag.Args()[0]

	err := confReader.GlobalConf.LoadConfigFromJSON(confFilePth)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := run(filePth, langType, casePth)

	switch strings.ToLower(outputMode) {
	case "pure":
		fmt.Println(result.Message)
	case "onlypssd":
		fmt.Println(result.PassMsg())
	case "json":
		text, err := result.DumpJSON()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(text)
		}
	}

	return
}
