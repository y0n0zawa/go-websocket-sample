package room

import (
	"context"
)

type key int

const idKey key = 0

func NewContext(ctx context.Context, roomID int) context.Context {
	return context.WithValue(ctx, idKey, roomID)
}

func FromContext(ctx context.Context) int {
	return ctx.Value(idKey).(int)
}
