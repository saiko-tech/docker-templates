package main

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainFunc(t *testing.T) {
	r, w, err := os.Pipe()
	require.NoError(t, err)

	originalStdout := os.Stdout
	defer func() { os.Stdout = originalStdout }()

	os.Stdout = w

	main()

	w.Close()

	output, err := io.ReadAll(r)
	require.NoError(t, err)

	assert.Equal(t, "hello world\n", string(output))
}
