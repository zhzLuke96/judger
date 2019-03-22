package complier

import (
	"../confReader"
	"../utils"
)

func ComplieCode(codeStr string, langType string) (fileName string, err error) {
	var cmdText string

	fileName, err = utils.SaveStrAsFile(codeStr)
	if err != nil {
		return fileName, err
	}

	cmdText, err = confReader.GlobalConf.GetComplieCmd(langType)
	if err != nil {
		panic(err)
	}

	_, err = utils.GetExecCmdOutput(cmdText+" ./"+fileName+".buf", "")
	if err != nil {
		return "", err
	}

	return fileName, nil
}
