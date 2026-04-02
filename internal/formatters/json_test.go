package formatters

import (
	"code/internal/diff"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatJSON(t *testing.T) {
	tree := []diff.DiffNode{
		{
			Key:   "follow",
			Kind:  diff.KindAdded,
			Right: false,
		},
	}

	got, err := formatJSON(tree)
	require.NoError(t, err)

	var decoded []map[string]any
	err = json.Unmarshal([]byte(got), &decoded)
	require.NoError(t, err)

	require.Len(t, decoded, 1)
	require.Equal(t, "follow", decoded[0]["key"])
	require.Equal(t, "added", decoded[0]["type"])
	require.Equal(t, false, decoded[0]["new_value"])
}
