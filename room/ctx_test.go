package room

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContext_FromContext(t *testing.T) {
	roomID := 1

	ctx := context.Background()

	ctx = NewContext(ctx, roomID)
	actual := FromContext(ctx)

	assert.Equal(t, roomID, actual, "The room is stored in the context")
}
