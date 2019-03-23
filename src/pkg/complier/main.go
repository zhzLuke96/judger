package complier

import (
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
		_, err = utils.GetExecCmdOutput(cmdText+" "+srcfilePth, "")
		if err != nil {
			return "", err
		}
	}

	fileNamePart, err = utils.GetFileNameFromPth(srcfilePth)
	if err != nil {
		return "", err
	}
	return fileNamePart, nil
}
