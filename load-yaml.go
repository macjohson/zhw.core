package zhwcore

import (
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

// 读取yaml文件
func LoadYaml(filePath string, v interface{}, logger *zap.Logger) (err error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0755)

	defer func() {
		closeErr := file.Close()

		if closeErr != nil {
			logger.Error("close yaml file fail", zap.Error(closeErr))
		}
	}()

	if err != nil {
		return err
	}

	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(v)

	return
}
