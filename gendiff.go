package code

import (
	"code/internal/diff"
	"code/internal/formatters"
	"code/internal/parser"
	"fmt"
)

// GenDiff вычисляет различия между двумя файлами и возвращает их
// в виде строки в указанном формате.
func GenDiff(leftPath, rightPath, format string) (string, error) {
	leftData, rightData, err := parseFiles(leftPath, rightPath)
	if err != nil {
		return "", err
	}

	rootDepth := 1
	tree := diff.BuildTree(leftData, rightData)

	return formatters.Format(tree, rootDepth, format)
}

func parseFiles(leftPath, rightPath string) (map[string]any, map[string]any, error) {
	leftData, err := parser.Parse(leftPath)
	if err != nil {
		return nil, nil, fmt.Errorf("parse left file: %w", err)
	}

	rightData, err := parser.Parse(rightPath)
	if err != nil {
		return nil, nil, fmt.Errorf("parse right file: %w", err)
	}

	return leftData, rightData, nil
}
