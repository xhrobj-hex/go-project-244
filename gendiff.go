package code

import (
	"code/internal/parser"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func GenDiff(filePathLeft, filePathRight, format string) (string, error) {
	_ = format // NOTE: пока не используем

	// распарсить файлы
	dataLeft, dataRight, err := parseFiles(filePathLeft, filePathRight)
	if err != nil {
		return "", err
	}

	// собрать все ключи из обоих объектов
	// отсортировать ключи
	keysSorted := sortedKeys(dataLeft, dataRight)

	// пройтись по ключам и собрать строки в diff (пока массив, но в итоге нужно будет дерево)
	diff := buildDiff(dataLeft, dataRight, keysSorted)

	// склеить все в итоговую строку
	r := fmt.Sprintf("{\n%s\n}", strings.Join(diff, "\n"))

	return r, nil
}

func parseFiles(filePath1, filePath2 string) (map[string]any, map[string]any, error) {
	fileData1, err := parser.Parse(filePath1)
	if err != nil {
		return nil, nil, err
	}

	fileData2, err := parser.Parse(filePath2)
	if err != nil {
		return nil, nil, err
	}

	return fileData1, fileData2, nil
}

func sortedKeys(data1, data2 map[string]any) []string {
	keysSet := make(map[string]struct{})
	for k := range data1 {
		keysSet[k] = struct{}{}
	}
	for k := range data2 {
		keysSet[k] = struct{}{}
	}

	keys := make([]string, 0, len(keysSet))
	for k := range keysSet {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
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
