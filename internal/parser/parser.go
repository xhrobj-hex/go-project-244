package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func Parse(filePath string) (map[string]any, error) {
	ext := strings.ToLower(filepath.Ext(filePath))

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data map[string]any // ???: а если массив в корне?

	switch ext {
	case ".json":
		err = json.Unmarshal(fileData, &data)
		if err != nil {
			return nil, err
		}
	case ".yml", ".yaml":
		err = yaml.Unmarshal(fileData, &data)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported file format: %s", ext)
	}

	return data, nil
}
