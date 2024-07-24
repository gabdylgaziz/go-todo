package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	_ "gotodo/docs"
	dto "gotodo/internal/domain/gotodo"
	"gotodo/internal/service/gotodo"
	"gotodo/pkg/server/response"
	"net/http"
	"time"
)

type TodoHandler struct {
	todoService *gotodo.Service
}

func NewTodoHandler(s *gotodo.Service) *TodoHandler {
	return &TodoHandler{todoService: s}
}

// @Summary     Health check
// @Description Healthcheking.
// @Tags        health
// @Accept      json
// @Produce     json
// @Success     200 {object} map[string]string
// @Router      /health [get]
func (h *TodoHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"Health": "OK!",
	}

	render.JSON(w, r, response)
}

func (h *TodoHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", h.CreateTask)
	r.Get("/", h.ListTasks)
	r.Put("/{id}", h.UpdateTask)
	r.Delete("/{id}", h.DeleteTask)
	r.Put("/{id}/done", h.MarkTaskAsDone)

	return r
}

// @Summary Create a new task
// @Description Create a new task with title and date
// @Accept json
// @Produce json
// @Param task body dto.TaskCreate true "Task creation request body"
// @Success 201 {object} dto.TaskResponse "Task created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 404 {object} map[string]interface{} "Task already exists"
// @Router /api/todo-list/tasks [post]
func (h *TodoHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	req := dto.TaskCreate{}
	if err := render.Bind(r, &req); err != nil {
		response.BadRequest(w, r, err, req)
		return
	}

	activeAt, err := time.Parse("2006-01-02", req.ActiveAt)
	if err != nil {
		response.BadRequest(w, r, err, req)
		return
	}

	task, err := h.todoService.CreateTask(r.Context(), req.Title, activeAt)
	if err != nil {
		response.NotFound(w, r, err)
		return
	}

	response.Created(w, r, task)
}

// @Summary Update an existing task
// @Description Update the title and date of an existing task
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param task body dto.TaskUpdate true "Task update request body"
// @Success 204 "Task updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 404 {object} map[string]interface{} "Task not found"
// @Router /api/todo-list/tasks/{id} [put]
func (h *TodoHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	req := dto.TaskUpdate{}
	if err := render.Bind(r, &req); err != nil {
		response.BadRequest(w, r, err, req)
		return
	}

	activeAt, err := time.Parse("2006-01-02", req.ActiveAt)
	if err != nil {
		response.BadRequest(w, r, err, req)
		return
	}

	err = h.todoService.UpdateTask(r.Context(), id, req.Title, activeAt)
	if err != nil {
		response.NotFound(w, r, err)
		return
	}

	response.NoContent(w, r)
}

// @Summary Delete a task
// @Description Delete a task by its ID
// @Param id path string true "Task ID"
// @Success 204 "Task deleted successfully"
// @Failure 404 {object} map[string]interface{} "Task not found"
// @Router /api/todo-list/tasks/{id} [delete]
func (h *TodoHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.todoService.DeleteTask(r.Context(), id)
	if err != nil {
		response.NotFound(w, r, err)
		return
	}

	response.NoContent(w, r)
}

// @Summary Mark a task as done
// @Description Mark a task as completed by its ID
// @Param id path string true "Task ID"
// @Success 204 "Task marked as done successfully"
// @Failure 404 {object} map[string]interface{} "Task not found"
// @Router /api/todo-list/tasks/{id}/done [put]
func (h *TodoHandler) MarkTaskAsDone(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.todoService.MarkTaskAsDone(r.Context(), id)
	if err != nil {
		response.NotFound(w, r, err)
		return
	}

	response.NoContent(w, r)
}

// @Summary List tasks by status
// @Description Retrieve tasks filtered by status (active or done)
// @Produce json
// @Param status query string false "Status of the tasks to retrieve (active or done)"
// @Success 200 {array} dto.TaskResponse "List of tasks"
// @Failure 400 {object} map[string]interface{} "Invalid status parameter"
// @Router /api/todo-list/tasks [get]
func (h *TodoHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	if status == "" {
		status = "active"
	}

	tasks := h.todoService.ListTasks(r.Context(), status)
	response.OK(w, r, tasks)
}
