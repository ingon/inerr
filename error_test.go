package inerr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWrap(t *testing.T) {
	err := New("abc:%d", 1)
	require.Error(t, err)
	require.Equal(t, "abc:1", err.Error())

	merr := Wrap(err, "cde:%d", 2)
	require.Error(t, merr)
	require.Equal(t, "cde:2: abc:1", merr.Error())
	require.ErrorIs(t, merr, err)
}
