package code

import (
	"code/internal/parser"
	"fmt"
	"sort"
	"strings"
)

func GenDiff(filepath1, filepath2, format string) (string, error) {
	// 1. распарсить первый файл
	json1, err := parser.Parse(filepath1)
	if err != nil {
		return "", err
	}

	// 2. распарсить второй файл
	json2, err := parser.Parse(filepath2)
	if err != nil {
		return "", err
	}

	// 3. собрать все ключи из обоих объектов
	keys := map[string]string{}
	for k := range json1 {
		keys[k] = ""
	}
	for k := range json2 {
		keys[k] = ""
	}

	// 4. отсортировать ключи
	keysSorted := []string{}
	for k := range keys {
		keysSorted = append(keysSorted, k)
	}
	sort.Strings(keysSorted)

	// 5. пройтись по ним и собрать строки diff
	diff := []string{}
	for _, key := range keysSorted {
		v1, ok := json1[key]
		if !ok {
			v2 := json2[key]
			d := fmt.Sprintf("  + %s: %v", key, v2)
			diff = append(diff, d)
			continue
		}
		v2, ok := json2[key]
		if !ok {
			d := fmt.Sprintf("  - %s: %v", key, v1)
			diff = append(diff, d)
			continue
		}
		if v1 != v2 {
			d := fmt.Sprintf("  - %s: %v", key, v1)
			diff = append(diff, d)
			d = fmt.Sprintf("  + %s: %v", key, v2)
			diff = append(diff, d)
			continue
		}
		d := fmt.Sprintf("    %s: %v", key, v1)
		diff = append(diff, d)
	}

	// 6. склеить все в итоговую строку
	r := fmt.Sprintf("{\n%s\n}", strings.Join(diff, "\n"))

	return r, nil
}
