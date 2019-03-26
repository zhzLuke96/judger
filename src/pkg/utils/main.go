package utils

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	Version = "a0.1"
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

func SaveStrAsFileMD5(content string) (filePth string, err error) {
	filePth = "./" + MD52File(content) + ".buf"
	return filePth, SaveStrAsFile(content, filePth)
}

func SaveStrAsFile(content string, filePth string) (err error) {
	var f *os.File
	var byteCount int

	if checkFileIsExist(filePth) {
		return nil
	} else {
		if f, err = os.Create(filePth); err != nil {
			return err
		} else {
			fmt.Printf("Code file named %s .\n", filePth)
		}
		defer f.Close()
	}
	w := bufio.NewWriter(f)
	defer w.Flush()

	if byteCount, err = w.WriteString(content); err != nil {
		return err
	}
	fmt.Printf("Writer in [%s] %d bytes\n", filePth, byteCount)
	return nil
}

func GetFileNameFromPth(filePth string) string {
	// var matchArr [][]string

	// reg, err := regexp.Compile(`(.*\/)?(.+)\..+`)
	// if err == nil {
	// 	matchArr = reg.FindAllStringSubmatch(filePth, -1)
	// 	if len(matchArr) != 0 {
	// 		return matchArr[0][2], nil
	// 	}
	// }

	// reg, err = regexp.Compile(`(.*\/)?(.+)`)
	// if err == nil {
	// 	matchArr = reg.FindAllStringSubmatch(filePth, -1)
	// 	if len(matchArr) != 0 {
	// 		return matchArr[0][2], nil
	// 	}
	// }

	// return "", fmt.Errorf("ERROR: '%s' is't file path, cant get filename", filePth)

	return strings.Replace(filepath.Base(filePth), filepath.Ext(filePth), "", 1)
}

func trimOutput(buffer bytes.Buffer) string {
	return strings.TrimSpace(string(bytes.TrimRight(buffer.Bytes(), "\x00")))
}

var ext2lang = map[string]string{
	".py":  "python",
	".js":  "javascript",
	".go":  "golang",
	".cpp": "cpp",
	".c":   "c"}

func FileNameToLang(filePth string) string {
	ext := filepath.Ext(filePth)
	// fmt.Printf("[LOG] ext = %v\n", ext)
	if v, ok := ext2lang[ext]; ok {
		return v
	}
	return ""
}
