package gotodo

import (
	"context"
	"gotodo/internal/domain/gotodo"
	"time"
)

func (s *Service) CreateTask(ctx context.Context, title string, activeAt time.Time) (res gotodo.TaskResponse, err error) {
	task, err := s.todoRepository.Add(ctx, title, activeAt)
	if err != nil {
		return
	}

	res = gotodo.ParseFromEntity(task)
	return
}

func (s *Service) ListTasks(ctx context.Context, status string) (res []gotodo.TaskResponse) {
	tasks := s.todoRepository.List(ctx, status)

	for _, task := range tasks {
		r := gotodo.ParseFromEntity(task)
		res = append(res, r)
	}
	return
}

func (s *Service) MarkTaskAsDone(ctx context.Context, id string) (err error) {
	err = s.todoRepository.MarkTaskAsDone(ctx, id)
	if err != nil {
		return
	}

	return
}

func (s *Service) UpdateTask(ctx context.Context, id string, title string, activeAt time.Time) (err error) {
	err = s.todoRepository.Update(ctx, id, title, activeAt)
	if err != nil {
		return
	}

	return
}

func (s *Service) DeleteTask(ctx context.Context, id string) (err error) {
	err = s.todoRepository.Delete(ctx, id)
	if err != nil {
		return
	}

	return
}
