package complier

import (
	"fmt"
	"strings"

	"../confReader"
	"../utils"
)

func ComplieCode(codeStr string, langType string) (fileName string, err error) {
	fileName, err = utils.SaveStrAsFileMD5AutoExtFromLangType(codeStr, langType)
	if err != nil {
		return fileName, err
	}

	return ComplieCodeFromFile(fileName, langType)
}

func ComplieCodeFromFile(srcfilePth string, langType string) (fileNamePart string, err error) {
	var cmdText string

	fileNamePart = utils.GetFileNameFromPth(srcfilePth)
	cmdText, err = confReader.GlobalConf.GetComplieCmd(langType)
	if err != nil {
		return "", err
	}
	cmdText = strings.Replace(cmdText, "<<SRCFILENAME>>", fileNamePart, -1)

	if cmdText != "" {
		// _, err = utils.GetExecCmdOutput(cmdText+" "+srcfilePth, "")
		pres, err := utils.ExecCmd(cmdText + " " + srcfilePth)
		if err != nil {
			return "", err
		}
		if pres.Stderr != "" {
			return "", fmt.Errorf(pres.Stderr)
		}
	}

	return fileNamePart, nil
}
