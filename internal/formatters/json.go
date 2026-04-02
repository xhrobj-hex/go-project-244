package formatters

import (
	"code/internal/diff"
	"encoding/json"
)

func formatJSON(tree []diff.DiffNode) (string, error) {
	data, err := json.Marshal(tree)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
