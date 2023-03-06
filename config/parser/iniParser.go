package parser

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type IniParser struct {
	filePath string
	file     *ini.File
}

func NewIniParser(filePath string) *IniParser {
	return &IniParser{
		filePath: filePath,
	}
}

func (i *IniParser) Parse() error {
	file, err := ini.Load(i.filePath)
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
		return nil
	}
	i.file = file
	return nil
}

func (i *IniParser) GetString(sec string, key string) string {
	return i.file.Section(sec).Key(key).String()
}
func (i *IniParser) GetInt(sec string, key string) int {
	v, _ := i.file.Section(sec).Key(key).Int()
	return v
}
