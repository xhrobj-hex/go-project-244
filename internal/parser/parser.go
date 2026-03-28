package parser

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Parse(path string) (map[string]any, error) {
	// 1. проверить расширение json
	ext := filepath.Ext(path)
	if strings.ToLower(ext) != ".json" {
		return nil, errors.New("file isn't json")
	}

	// 2. открыть файл
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = file.Close()
	}()

	// 3. прочитать содержимое
	filedata, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// 4. распарсить в json
	var jsondata map[string]any // ???: а если массив в корне?
	err = json.Unmarshal(filedata, &jsondata)
	if err != nil {
		return nil, err
	}

	return jsondata, nil
}
