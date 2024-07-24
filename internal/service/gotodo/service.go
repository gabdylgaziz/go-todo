package gotodo

import "gotodo/internal/domain/gotodo"

type Configuration func(s *Service) error

type Service struct {
	todoRepository gotodo.Repository
}

func New(configs ...Configuration) (s *Service, err error) {
	s = &Service{}

	for _, cfg := range configs {
		if err = cfg(s); err != nil {
			return
		}
	}
	return
}

func WithTodoRepository(todoRepository gotodo.Repository) Configuration {
	return func(s *Service) error {
		s.todoRepository = todoRepository
		return nil
	}
}
