package parser

import (
	"fmt"
	yamlConfig "github.com/olebedev/config"
	"strings"
)

type YamlParser struct {
	filePath string
	file     *yamlConfig.Config
}

func NewYamlParser(filePath string) *YamlParser {
	return &YamlParser{
		filePath: filePath,
	}
}

func (i *YamlParser) Parse() error {
	file, err := yamlConfig.ParseYamlFile(i.filePath)
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
		return nil
	}
	//fmt.Println("file", file)
	i.file = file
	return nil
}

func (i *YamlParser) GetString(keys ...string) string {
	key := strings.Join(keys, ".")
	//fmt.Println("GetString....")
	//fmt.Println(key)
	val, _ := i.file.String(key)
	return val
}
func (i *YamlParser) GetInt(keys ...string) int {
	key := strings.Join(keys, ".")
	val, _ := i.file.Int(key)
	return val
}
