package archive

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.podman.io/image/v5/types"
)

func TestDesiredLayerCompression(t *testing.T) {
	preserve := types.PreserveOriginal
	compress := types.Compress
	decompress := types.Decompress

	tests := []struct {
		name     string
		sysCtx   *types.SystemContext
		expected types.LayerCompression
	}{
		{
			name:     "nil SystemContext defaults to PreserveOriginal",
			sysCtx:   nil,
			expected: types.PreserveOriginal,
		},
		{
			name:     "nil DockerArchiveLayerCompression defaults to PreserveOriginal",
			sysCtx:   &types.SystemContext{},
			expected: types.PreserveOriginal,
		},
		{
			name:     "explicit PreserveOriginal",
			sysCtx:   &types.SystemContext{DockerArchiveLayerCompression: &preserve},
			expected: types.PreserveOriginal,
		},
		{
			name:     "explicit Compress",
			sysCtx:   &types.SystemContext{DockerArchiveLayerCompression: &compress},
			expected: types.Compress,
		},
		{
			name:     "explicit Decompress",
			sysCtx:   &types.SystemContext{DockerArchiveLayerCompression: &decompress},
			expected: types.Decompress,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &archiveImageDestination{
				sysCtx: tt.sysCtx,
			}
			assert.Equal(t, tt.expected, d.DesiredLayerCompression())
		})
	}
}
