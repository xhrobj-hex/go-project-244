package diff

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildTree(t *testing.T) {
	left := map[string]any{
		"follow":  false,
		"host":    "hexlet.io",
		"proxy":   "123.234.53.22",
		"timeout": 50,
		"nest": map[string]any{
			"key": "value",
		},
	}

	right := map[string]any{
		"follow":  false,
		"host":    "hexlet.io",
		"timeout": 20,
		"verbose": true,
		"nest": map[string]any{
			"key":  "value",
			"key2": "value2",
		},
	}

	got := BuildTree(left, right)

	require.Len(t, got, 6)

	require.Equal(t, "follow", got[0].Key)
	require.Equal(t, KindUnchanged, got[0].Kind)
	require.Equal(t, false, got[0].Left)
	require.Equal(t, false, got[0].Right)

	require.Equal(t, "host", got[1].Key)
	require.Equal(t, KindUnchanged, got[1].Kind)
	require.Equal(t, "hexlet.io", got[1].Left)
	require.Equal(t, "hexlet.io", got[1].Right)

	require.Equal(t, "nest", got[2].Key)
	require.Equal(t, KindNested, got[2].Kind)
	require.Len(t, got[2].Children, 2)

	require.Equal(t, "key", got[2].Children[0].Key)
	require.Equal(t, KindUnchanged, got[2].Children[0].Kind)

	require.Equal(t, "key2", got[2].Children[1].Key)
	require.Equal(t, KindAdded, got[2].Children[1].Kind)
	require.Equal(t, "value2", got[2].Children[1].Right)

	require.Equal(t, "proxy", got[3].Key)
	require.Equal(t, KindRemoved, got[3].Kind)
	require.Equal(t, "123.234.53.22", got[3].Left)

	require.Equal(t, "timeout", got[4].Key)
	require.Equal(t, KindChanged, got[4].Kind)
	require.Equal(t, 50, got[4].Left)
	require.Equal(t, 20, got[4].Right)

	require.Equal(t, "verbose", got[5].Key)
	require.Equal(t, KindAdded, got[5].Kind)
	require.Equal(t, true, got[5].Right)
}
