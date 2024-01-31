package zhwcore

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// 读取yaml文件
func LoadYaml(filePath string, v interface{}) (err error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0755)

	defer func() {
		closeErr := file.Close()

		if closeErr != nil {
			fmt.Println(closeErr)
		}
	}()

	if err != nil {
		return err
	}

	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(v)

	return
}
