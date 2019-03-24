package confReader

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"../utils"
)

var GlobalConf = *CreateConf()

// complier config struct, Store compiler call instructions
type jConfig struct {
	conf   map[string]map[string]string
	loaded bool
}

// init function, loaded should be false.
func CreateConf() (newConfig *jConfig) {
	newConfig = new(jConfig)
	newConfig.loaded = false
	return newConfig
}

// Read compiler call instructions from JSON format configuration files
func (c *jConfig) LoadConfigFromJSON(fileName string) (err error) {
	configContent, err := utils.ReadFile(fileName)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(configContent, &c.conf); err != nil {
		return err
	}
	c.loaded = true
	return nil
}

// Determine whether the configuration object is
// available and report the corresponding error
func (c *jConfig) keyTesting(key string) (err error) {
	if _, ok := c.conf[key]; !ok {
		if c.loaded {
			return fmt.Errorf("Language [%s] config not found", key)
		}
		return errors.New("Config object is empty or not loaded")
	}
	return nil
}

// Gets the call instructions of the compiler in the specified language
func (c *jConfig) GetComplieCmd(langType string) (cmdContext string, err error) {
	if err = c.keyTesting(langType); err != nil {
		return "", err
	}
	return c.conf[langType]["complie"], nil
}

// Get program call instructions in the specified language
func (c *jConfig) GetRunCmd(langType string) (cmdContext string, err error) {
	if err = c.keyTesting(langType); err != nil {
		return "", err
	}
	return c.conf[langType]["run"] + " " + c.conf[langType]["pth"], nil
}

// Get program call instructions with actual program file name
func (c *jConfig) GetRunCmdWithActualFileName(langType string, fileNamePart string) (cmdContext string, err error) {
	var commenCmd string

	commenCmd, err = c.GetRunCmd(langType)
	return strings.Replace(commenCmd, "<<SRCFILENAME>>", fileNamePart, 1), err
}
