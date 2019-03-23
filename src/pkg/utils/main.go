package utils

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func ReadFile(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func MD52File(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filePth string) bool {
	var exist = true
	if _, err := os.Stat(filePth); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func SaveStrAsFile(content string) (filePth string, err error) {
	var f *os.File
	var byteCount int

	filePth = "./" + MD52File(content) + ".buf"
	if checkFileIsExist(filePth) {
		return filePth, nil
	} else {
		if f, err = os.Create(filePth); err != nil {
			return filePth, err
		} else {
			fmt.Printf("Code file named %s .\n", filePth)
		}
		defer f.Close()
	}
	w := bufio.NewWriter(f)
	defer w.Flush()

	if byteCount, err = w.WriteString(content); err != nil {
		return filePth, err
	}
	fmt.Printf("Writer in [%s] %d bytes\n", filePth, byteCount)
	return filePth, nil
}

func GetExecCmdOutput(cmdcontent string, stdin string) (output string, err error) {
	var out []byte
	var stderr bytes.Buffer

	args := strings.Fields(cmdcontent)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = &stderr

	if stdin != "" {
		cmd.Stdin = strings.NewReader(stdin)
	}

	out, err = cmd.Output()
	if len(stderr.String()) != 0 {
		return "", errors.New(stderr.String())
	}
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func GetFileNameFromPth(filePth string) (fileName string, err error) {
	var matchArr [][]string

	reg, err := regexp.Compile(`(.*\/)?(.+)\..+`)
	if err == nil {
		matchArr = reg.FindAllStringSubmatch(filePth, -1)
		if len(matchArr) != 0 {
			return matchArr[0][2], nil
		}
	}

	reg, err = regexp.Compile(`(.*\/)?(.+)`)
	if err == nil {
		matchArr = reg.FindAllStringSubmatch(filePth, -1)
		if len(matchArr) != 0 {
			return matchArr[0][2], nil
		}
	}

	return "", fmt.Errorf("ERROR: '%s' is't file path, cant get filename", filePth)
}
