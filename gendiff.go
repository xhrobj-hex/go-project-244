package code

import (
	"code/internal/parser"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func GenDiff(filePath1, filepath2, format string) (string, error) {
	_ = format // NOTE: пока не используем

	// распарсить файлы
	data1, data2, err := parseFiles(filePath1, filepath2)
	if err != nil {
		return "", err
	}

	// собрать все ключи из обоих объектов
	keys := map[string]struct{}{}
	for k := range data1 {
		keys[k] = struct{}{}
	}
	for k := range data2 {
		keys[k] = struct{}{}
	}

	// отсортировать ключи
	keysSorted := make([]string, 0, len(keys))
	for k := range keys {
		keysSorted = append(keysSorted, k)
	}
	sort.Strings(keysSorted)

	// пройтись по ключам и собрать строки в diff (пока массив, но в итоге нужно будет дерево)
	diff := buildDiff(data1, data2, keysSorted)

	// склеить все в итоговую строку
	r := fmt.Sprintf("{\n%s\n}", strings.Join(diff, "\n"))

	return r, nil
}

func parseFiles(filePath1, filePath2 string) (map[string]any, map[string]any, error) {
	// 1. распарсить первый файл
	fileData1, err := parser.Parse(filePath1)
	if err != nil {
		return nil, nil, err
	}

	// 2. распарсить второй файл
	fileData2, err := parser.Parse(filePath2)
	if err != nil {
		return nil, nil, err
	}

	return fileData1, fileData2, nil
}

func buildDiff(data1, data2 map[string]any, keys []string) []string {
	diff := make([]string, 0)

	for _, key := range keys {
		v1, ok1 := data1[key]
		v2, ok2 := data2[key]

		switch {
		case !ok1:
			d := fmt.Sprintf("  + %s: %v", key, v2)
			diff = append(diff, d)
		case !ok2:
			d := fmt.Sprintf("  - %s: %v", key, v1)
			diff = append(diff, d)
		case !reflect.DeepEqual(v1, v2):
			d := fmt.Sprintf("  - %s: %v", key, v1)
			diff = append(diff, d)

			d = fmt.Sprintf("  + %s: %v", key, v2)
			diff = append(diff, d)
		default:
			d := fmt.Sprintf("    %s: %v", key, v1)
			diff = append(diff, d)
		}
	}

	return diff
}
