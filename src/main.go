package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"./pkg/RPC"
	"./pkg/confReader"
	"./pkg/retMsg"
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

func main() {
	var (
		langType   string
		casePth    string
		filePth    string
		outputMode string
		onlog      bool
		RPCSev     bool
		sevport    string
	)
	flag.StringVar(&langType, "lang", "auto", "Target language type of input code.")
	flag.StringVar(&langType, "l", "auto", "Target language type of input code.")

	flag.StringVar(&casePth, "case", "0x1", "the file path of the test case.")
	flag.StringVar(&casePth, "c", "0x1", "the file path of the test case.")

	flag.StringVar(&outputMode, "mode", "json", "choice program output style [pure,json,onlyPass].")
	flag.StringVar(&outputMode, "m", "json", "choice program output style [pure,json,onlyPass].")

	flag.BoolVar(&onlog, "log", false, "write log in file.")

	flag.BoolVar(&RPCSev, "sev", false, "runing a RPC server for judger.")
	flag.StringVar(&sevport, "port", "8088", "RPC server port.")

	flag.Parse()

	// Loading config
	err := confReader.GlobalConf.LoadConfigFromJSON(confFilePth)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ========================RPC
	if RPCSev {
		code := RPC.RunLocalRPCSev(sevport)
		os.Exit(code)
	}
	// RPC========================

	if len(flag.Args()) == 0 {
		fmt.Println("flag needs an argument for fileName.")
		flag.Usage()
		return
	}
	filePth = flag.Args()[0]

	if langType == "auto" {
		langType = utils.FileNameToLang(filePth)
		if langType == "" {
			fmt.Println("The current language pattern is automatic and cannot be recognized correctly. Please set up the correct programming language and try again.")
			flag.Usage()
			return
		}
	}
	result := retMsg.Run(filePth, langType, casePth)

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
