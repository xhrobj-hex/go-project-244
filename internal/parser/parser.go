package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Parse(path string) (map[string]any, error) {
	ext := strings.ToLower(filepath.Ext(path))
	if ext != ".json" {
		return nil, fmt.Errorf("unsupported file format: %s", ext)
	}

	fileData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data map[string]any // ???: а если массив в корне?
	if err := json.Unmarshal(fileData, &data); err != nil {
		return nil, err
	}

	return data, nil
}
