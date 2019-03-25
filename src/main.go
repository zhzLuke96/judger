package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"./pkg/complier"
	"./pkg/confReader"
	"./pkg/ojtest"
	"./pkg/utils"
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
		onlog      bool
	)
	flag.StringVar(&langType, "lang", "auto", "Target language type of input code.")
	flag.StringVar(&langType, "l", "auto", "Target language type of input code.")

	flag.StringVar(&casePth, "case", "0x1", "the file path of the test case.")
	flag.StringVar(&casePth, "c", "0x1", "the file path of the test case.")

	flag.StringVar(&outputMode, "mode", "json", "choice program output style [pure,json,onlyPass].")
	flag.StringVar(&outputMode, "m", "json", "choice program output style [pure,json,onlyPass].")

	flag.BoolVar(&onlog, "log", false, "witer log in file.")

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

	if langType == "auto" {
		langType = utils.FileNameToLang(filePth)
		if langType == "" {
			fmt.Println("The current language pattern is automatic and cannot be recognized correctly. Please set up the correct programming language and try again.")
			flag.Usage()
			return
		}
	}
	result := run(filePth, langType, casePth)

	var outputContent string

	switch strings.ToLower(outputMode) {
	case "pure":
		outputContent = result.Message
	case "onlypssd":
		outputContent = result.PassMsg()
	case "json":
		text, err := result.DumpJSON()
		if err != nil {
			outputContent = err.Error()
		} else {
			outputContent = text
		}
	}
	fmt.Println(outputContent)
	if onlog {
		utils.SaveStrAsFile(outputContent, "judger.log")
	}
	return
}
