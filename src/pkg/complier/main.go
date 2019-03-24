package complier

import (
	"fmt"

	"../confReader"
	"../utils"
)

func ComplieCode(codeStr string, langType string) (fileName string, err error) {
	fileName, err = utils.SaveStrAsFile(codeStr)
	if err != nil {
		return fileName, err
	}

	return ComplieCodeFromFile(fileName, langType)
}

func ComplieCodeFromFile(srcfilePth string, langType string) (fileNamePart string, err error) {
	var cmdText string

	cmdText, err = confReader.GlobalConf.GetComplieCmd(langType)
	if err != nil {
		return "", err
	}

	if cmdText != "" {
		// _, err = utils.GetExecCmdOutput(cmdText+" "+srcfilePth, "")
		_, stderr, err := utils.ShellCmd(cmdText+" "+srcfilePth, "")
		if err != nil {
			return "", err
		}
		if stderr != "" {
			return "", fmt.Errorf(stderr)
		}
	}

	fileNamePart, err = utils.GetFileNameFromPth(srcfilePth)
	if err != nil {
		return "", err
	}
	return fileNamePart, nil
}
