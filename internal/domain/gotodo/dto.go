package gotodo

import (
	"errors"
	"net/http"
)

type TaskCreate struct {
	Title    string `json:"title"`
	ActiveAt string `json:"active_at"`
}

func (s *TaskCreate) Bind(r *http.Request) error {
	if s.Title == "" {
		return errors.New("title: cannot be blank")
	}

	if len(s.Title) > 200 {
		return errors.New("title: title too long")
	}

	if s.ActiveAt == "" {
		return errors.New("active_at: cannot be blank")
	}

	return nil
}

type TaskUpdate struct {
	Title    string `json:"title"`
	ActiveAt string `json:"active_at"`
}

func (s *TaskUpdate) Bind(r *http.Request) error {
	if s.Title == "" {
		return errors.New("title: cannot be blank")
	}

	if s.ActiveAt == "" {
		return errors.New("active_at: cannot be blank")
	}

	return nil
}

type TaskResponse struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	ActiveAt string `json:"active_at"`
	Done     bool   `json:"done"`
}

func ParseFromEntity(data Entity) (res TaskResponse) {
	activeAt := data.ActiveAt.Format("2006-01-02")
	res = TaskResponse{
		ID:       data.ID,
		Title:    data.Title,
		ActiveAt: activeAt,
		Done:     data.Done,
	}
	return
}
