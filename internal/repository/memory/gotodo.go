package memory

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gotodo/internal/domain/gotodo"
	"sync"
	"time"
)

type TodoRepository struct {
	tasks map[string]gotodo.Entity
	mu    sync.Mutex
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		tasks: make(map[string]gotodo.Entity),
	}
}

func (r *TodoRepository) Add(ctx context.Context, title string, activeAt time.Time) (task gotodo.Entity, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, task := range r.tasks {
		if task.Title == title && task.ActiveAt == activeAt {
			return gotodo.Entity{}, errors.New("task already exists")
		}
	}

	id := uuid.New().String()
	task = gotodo.Entity{
		ID:       id,
		Title:    title,
		ActiveAt: activeAt,
		Done:     false,
	}

	r.tasks[id] = task
	return task, nil
}

func (r *TodoRepository) Get(ctx context.Context, id string) (task gotodo.Entity, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	task, exists := r.tasks[id]
	if !exists {
		return gotodo.Entity{}, errors.New("task not found")
	}
	return task, nil
}

func (r *TodoRepository) Update(ctx context.Context, id, title string, activeAt time.Time) (err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	task, exists := r.tasks[id]
	if !exists {
		return errors.New("task not found")
	}

	task.Title = title
	task.ActiveAt = activeAt
	r.tasks[id] = task
	return nil
}

func (r *TodoRepository) Delete(ctx context.Context, id string) (err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(r.tasks, id)
	return nil
}

func (r *TodoRepository) MarkTaskAsDone(ctx context.Context, id string) (err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	task, exists := r.tasks[id]
	if !exists {
		return errors.New("task not found")
	}

	task.Done = true
	r.tasks[id] = task
	return nil
}

func (r *TodoRepository) List(ctx context.Context, status string) (tasks []gotodo.Entity) {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	for _, task := range r.tasks {
		if status == "active" && !task.Done && task.ActiveAt.Before(now) {
			tasks = append(tasks, task)
		} else if status == "done" && task.Done {
			tasks = append(tasks, task)
		}
	}
	return tasks
}
