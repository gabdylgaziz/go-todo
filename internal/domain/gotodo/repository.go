package gotodo

import (
	"context"
	"time"
)

type Repository interface {
	Add(ctx context.Context, title string, activeAt time.Time) (task Entity, err error)
	Get(ctx context.Context, id string) (task Entity, err error)
	Update(ctx context.Context, id, title string, activeAt time.Time) (err error)
	Delete(ctx context.Context, id string) (err error)
	MarkTaskAsDone(ctx context.Context, id string) (err error)
	List(ctx context.Context, status string) (tasks []Entity)
}
