package config

import (
	"bufio"
	"os"
	"regexp"
    "fmt"
)

type ConfigFile struct {
	filename  string
	separator string
	m         map[string]string
}

func Config(filename string, separator string) (cf *ConfigFile, err error) {
	return &ConfigFile{filename, separator, make(map[string]string)}, nil
}

func (cf *ConfigFile) Parse() (* map[string]string, error) {
	f, err := os.Open(cf.filename)
	if err != nil {
		return nil, err
	}

    defer f.Close()

	reader := bufio.NewReader(f)
	restr := "^\\s*([a-zA-Z_][0-9a-zA-Z_]*)\\s*" + cf.separator + "\\s*([^\\s]+)\\s*$"
	re, err := regexp.Compile(restr)
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
	return &cf.m, err
}

func (cf *ConfigFile) Get(key string) string {
	return cf.m[key]
}

func (cf*ConfigFile) Set(key string, val string) {
   cf.m[key] = val; 
}

func (cf *ConfigFile) Save() {
    f, err := os.OpenFile(cf.filename, os.O_WRONLY, 0600);
    if (err != nil) {
        panic(err);
    }
    defer f.Close()
    for k,v := range cf.m {
        s := fmt.Sprintf("%s"+cf.separator+ "%s\n", k, v);
        f.WriteString(s)
    }
}

func (cf *ConfigFile) Close() {
}
