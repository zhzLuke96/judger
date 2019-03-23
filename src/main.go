package main

import (
	"flag"
	"fmt"
	"os"

	"./pkg/complier"
	"./pkg/confReader"
	"./pkg/ojtest"
)

const confFilePth = "./conf.json"

func clearCompliceFile() (err error) {
	os.RemoveAll("./__pycache__/")
	return nil
}

func run(codeFilePth string, langType string, casePth string) (result string) {
	programFilePth, err := complier.ComplieCodeFromFile(codeFilePth, langType)
	if err != nil {
		return err.Error()
	}

	per100, err := ojtest.RunTests(programFilePth, casePth, langType)
	if err != nil {
		return err.Error() + fmt.Sprintf(".\n\nPassed %d%%", per100)
	}

	return fmt.Sprintf("Passed %d%%.", per100)
}

func main() {
	var (
		langType string
		casePth  string
		filePth  string
	)
	flag.StringVar(&langType, "lang", "cpp", "Target language type of input code.")
	flag.StringVar(&langType, "l", "cpp", "Target language type of input code.")

	flag.StringVar(&casePth, "case", "0x1", "the file path of the test case.")
	flag.StringVar(&casePth, "c", "0x1", "the file path of the test case.")

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
	fmt.Println(result)
	return

	// result := run("../adder_case/adder.py", "py", "../adder_case/adder_testcase.json")
	// fmt.Println(result)

	// fmt.Printf("langType = %s\n", langType)
	// fmt.Printf("casePth = %s\n", casePth)
	// fmt.Printf("filePth = %s\n", filePth)

	// fmt.Println("THX 4 USEING!")
}
