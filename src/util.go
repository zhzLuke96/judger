package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func ReadAll(filePth string) ([]byte, error) {
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
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func saveStrAsFile(content string) (fileName string, err error) {
	var f *os.File
	var byteCount int

	fileName = MD52File(content)
	allName := "./" + fileName + ".buf"
	if checkFileIsExist(allName) {
		return fileName, nil
	} else {
		if f, err = os.Create(allName); err != nil {
			return fileName, err
		} else {
			fmt.Printf("Code file named %s .\n", allName)
		}
		defer f.Close()
	}
	w := bufio.NewWriter(f)
	defer w.Flush()

	if byteCount, err = w.WriteString(content); err != nil {
		return fileName, err
	}
	fmt.Printf("Writer in [%s] %d bytes\n", allName, byteCount)
	return fileName, nil
}

func getExecCmdOutput(cmdcontent string, stdin string) (output string, err error) {
	var out []byte
	args := strings.Fields(cmdcontent)
	cmd := exec.Command(args[0], args[1:]...)

	if stdin != "" {
		cmd.Stdin = strings.NewReader(stdin)
	}

	out, err = cmd.Output()
	return string(out[:]), err
}
