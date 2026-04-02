package formatters

import (
	"code/internal/diff"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatUnknownFormat(t *testing.T) {
	tree := []diff.DiffNode{}

	_, err := Format(tree, 1, "unknown")

	require.Error(t, err)
	require.ErrorContains(t, err, "(о_0) unknown output format")
}

func TestFormatStylishOption(t *testing.T) {
	tree := []diff.DiffNode{}
	_, err := Format(tree, 1, "stylish")
	require.NoError(t, err)
}

func TestFormatPlainOption(t *testing.T) {
	tree := []diff.DiffNode{}
	_, err := Format(tree, 1, "plain")
	require.NoError(t, err)
}

func TestFormatJSONOption(t *testing.T) {
	tree := []diff.DiffNode{}
	_, err := Format(tree, 1, "json")
	require.NoError(t, err)
}

func TestFormatEmptyFormatUsesStylish(t *testing.T) {
	tree := []diff.DiffNode{}

	got, err := Format(tree, 1, "")
	require.NoError(t, err)
	require.Equal(t, "{\n}", got)
}
