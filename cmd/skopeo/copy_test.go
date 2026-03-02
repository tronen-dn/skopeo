package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	// Invalid command-line arguments
	for _, args := range [][]string{
		{},
		{"a1"},
		{"a1", "a2", "a3"},
	} {
		out, err := runSkopeo(append([]string{"--insecure-policy", "copy"}, args...)...)
		assertTestFailed(t, out, err, "Exactly two arguments expected")
	}

	// FIXME: Much more test coverage
	// Actual feature tests exist in integration and systemtest
}

func TestDockerArchiveLayerCompressionFlag(t *testing.T) {
	// Invalid value
	_, err := runSkopeo("--insecure-policy", "copy",
		"--docker-archive-layer-compression", "invalid",
		"docker://localhost/nonesuch:latest", "docker-archive:/dev/null")
	assert.ErrorContains(t, err, `invalid --docker-archive-layer-compression value "invalid"`)

	// Valid values should be accepted (will fail later for other reasons, not flag parsing)
	for _, val := range []string{"preserve", "compress", "decompress"} {
		_, err := runSkopeo("--insecure-policy", "copy",
			"--docker-archive-layer-compression", val,
			"docker://localhost/nonesuch:latest", "docker-archive:/dev/null")
		if err != nil {
			assert.NotContains(t, err.Error(), "invalid --docker-archive-layer-compression")
		}
	}
}
