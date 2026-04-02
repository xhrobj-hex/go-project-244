package formatter

import (
	"code/internal/diff"
	"fmt"
)

// FormatTree форматирует дерево различий в строку согласно указанному формату.
func FormatTree(tree []diff.DiffNode, depth int, format string) (string, error) {
	switch format {
	case "stylish":
		return formatTreeStylish(tree, depth)
	default:
		return "", fmt.Errorf("(о_0) unknown format: %q", format)
	}
}
