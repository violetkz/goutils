package config

import (
	"bufio"
	"os"
	"regexp"
)

type ConfigFile struct {
	filename  string
	separator string
	file      *os.File
	m         map[string]string
}

func Config(filename string, separator string) (cf *ConfigFile, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return &ConfigFile{filename, separator, f, make(map[string]string)}, nil
}

func (cf *ConfigFile) Parse() (map[string]string, error) {

	reader := bufio.NewReader(cf.file)
	re, err := regexp.Compile("^\\s*([a-zA-Z_][0-9a-zA-Z_]*)\\s*" + cf.separator + "\\s*([^\\s]+)\\s*$")
	if err != nil {
		return cf.m, err
	}

	for {
		line, isPrefix, err := reader.ReadLine()
		if err == nil && isPrefix == false {
			subs := re.FindSubmatch(line)

			if len(subs) == 3 {
				str := string(subs[1])
				cf.m[str] = string(subs[2])
			}

		} else {
			break
		}
	}
	return cf.m, err
}

func (cf *ConfigFile) Get(key string) string {
	return cf.m[key]
}

