package formatters

import (
	"code/internal/diff"
	"fmt"
)

// Format форматирует дерево различий в строку согласно указанному формату.
func Format(tree []diff.DiffNode, depth int, format string) (string, error) {
	switch format {
	case "", "stylish":
		return formatStylish(tree, depth)
	case "plain":
		return formatPlain(tree)
	case "json":
		return formatJSON(tree)
	default:
		return "", fmt.Errorf("(о_0) unknown output format: %q", format)
	}
}
