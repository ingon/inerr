package inerr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrint(t *testing.T) {
	err := New("abc:%d", 1)
	require.Equal(t, "abc:1", Sprint(err))

	merr := Wrap(err, "cde:%d", 2)
	require.Equal(t, "cde:2\n  abc:1", Sprint(merr))

	kerr := Wrap(merr, "efg:%d", 3)
	require.Equal(t, "efg:3\n  cde:2\n    abc:1", Sprint(kerr))
}
