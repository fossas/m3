package tsz

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadBits(t *testing.T) {
	byteStream := []byte{
		0xca, 0xfe, 0xfd, 0x89, 0x1a, 0x2b, 0x3c, 0x48, 0x55, 0xe6, 0xf7,
		0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80,
	}
	is := newIStream(bytes.NewReader(byteStream))
	numBits := []int{1, 3, 4, 8, 7, 2, 64, 64}
	var res []uint64
	for _, v := range numBits {
		read, err := is.ReadBits(v)
		require.NoError(t, err)
		res = append(res, read)
	}
	expected := []uint64{0x1, 0x4, 0xa, 0xfe, 0x7e, 0x3, 0x1234567890abcdef, 0x1}
	require.Equal(t, expected, res)
	require.NoError(t, is.err)

	_, err := is.ReadBits(8)
	require.Error(t, err)
	require.Error(t, is.err)
}

func TestPeekBitsSuccess(t *testing.T) {
	byteStream := []byte{0xa9, 0xfe, 0xfe, 0xdf, 0x9b, 0x57, 0x21, 0xf1}
	is := newIStream(bytes.NewReader(byteStream))
	inputs := []struct {
		numBits  int
		expected uint64
	}{
		{0, 0},
		{1, 0x1},
		{8, 0xa9},
		{10, 0x2a7},
		{13, 0x153f},
		{16, 0xa9fe},
		{32, 0xa9fefedf},
		{64, 0xa9fefedf9b5721f1},
	}
	for _, input := range inputs {
		res, err := is.PeekBits(input.numBits)
		require.NoError(t, err)
		require.Equal(t, input.expected, res)
	}
	require.NoError(t, is.err)
	require.Equal(t, byte(0), is.current)
	require.Equal(t, 0, is.remaining)
}

func TestPeekBitsError(t *testing.T) {
	byteStream := []byte{0x1, 0x2}
	is := newIStream(bytes.NewReader(byteStream))
	res, err := is.PeekBits(20)
	require.Error(t, err)
	require.Equal(t, uint64(0), res)
}

func TestReadAfterPeekBits(t *testing.T) {
	byteStream := []byte{0xab, 0xcd}
	is := newIStream(bytes.NewReader(byteStream))
	res, err := is.PeekBits(10)
	require.NoError(t, err)
	require.Equal(t, uint64(0x2af), res)
	res, err = is.PeekBits(20)
	require.Error(t, err)

	inputs := []struct {
		numBits  int
		expected uint64
	}{
		{2, 0x2},
		{9, 0x15e},
	}
	for _, input := range inputs {
		res, err := is.ReadBits(input.numBits)
		require.NoError(t, err)
		require.Equal(t, input.expected, res)
	}
	res, err = is.ReadBits(8)
	require.Error(t, err)
}

func TestResetIStream(t *testing.T) {
	is := newIStream(bytes.NewReader(nil))
	is.ReadBits(1)
	require.Error(t, is.err)
	is.Reset(bytes.NewReader(nil))
	require.NoError(t, is.err)
	require.Equal(t, byte(0), is.current)
	require.Equal(t, 0, is.remaining)
}
