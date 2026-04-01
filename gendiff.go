package code

import (
	"code/internal/parser"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func GenDiff(filepath1, filepath2, format string) (string, error) {
	_ = format // NOTE: пока не используем

	// 1. распарсить первый файл
	data1, err := parser.Parse(filepath1)
	if err != nil {
		return "", err
	}

	// 2. распарсить второй файл
	data2, err := parser.Parse(filepath2)
	if err != nil {
		return "", err
	}

	// 3. собрать все ключи из обоих объектов
	keys := map[string]string{}
	for k := range data1 {
		keys[k] = ""
	}
	for k := range data2 {
		keys[k] = ""
	}

	// 4. отсортировать ключи
	keysSorted := []string{}
	for k := range keys {
		keysSorted = append(keysSorted, k)
	}
	sort.Strings(keysSorted)

	// 5. пройтись по ключам и собрать строки diff
	diff := make([]string, 0)

	for _, key := range keysSorted {
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

	// 6. склеить все в итоговую строку
	r := fmt.Sprintf("{\n%s\n}", strings.Join(diff, "\n"))

	return r, nil
}
